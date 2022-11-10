package neworder

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/fixt11"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/tag"

	"sylr.dev/fix/config"
	"sylr.dev/fix/pkg/cli/complete"
	"sylr.dev/fix/pkg/dict"
	"sylr.dev/fix/pkg/errors"
	"sylr.dev/fix/pkg/initiator"
	"sylr.dev/fix/pkg/initiator/application"
	"sylr.dev/fix/pkg/utils"
)

var (
	optionOrderSide, optionOrderType string
	optionOrderSymbol, optionOrderID string
	optionOrderExpiry                string
	optionOrderQuantity              int64
	optionOrderPrice                 float64
	optionOrderOrigination           string
	optionOrderAttributeType         []string
	optionOrderAttributeValue        []string
	optionPartyIDs                   []string
	optionPartyIDSources             []string
	optionPartySubIDs                []string
	optionPartySubIDTypes            []string
	optionPartyRoles                 []string
	optionPartyRoleQualifiers        []string
	optionExecReports                int
	optionExecReportsTimeout         time.Duration
	optionExecReportsTimeoutReset    bool
)

var NewOrderCmd = &cobra.Command{
	Use:               "order",
	Short:             "New single order",
	Long:              "Send a new single order after initiating a session with a FIX acceptor.",
	Args:              cobra.ExactArgs(0),
	ValidArgsFunction: cobra.NoFileCompletions,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if err := utils.ValidateRequiredFlags(cmd); err != nil {
			return err
		}

		if err := Validate(cmd, args); err != nil {
			return err
		}

		if cmd.HasParent() {
			parent := cmd.Parent()
			if parent.PersistentPreRunE != nil {
				return parent.PersistentPreRunE(cmd, args)
			}
		}
		return nil
	},
	RunE: Execute,
}

func init() {
	NewOrderCmd.Flags().StringVar(&optionOrderSide, "side", "", "Order side (buy, sell ... etc)")
	NewOrderCmd.Flags().StringVar(&optionOrderType, "type", "", "Order type (market, limit, stop ... etc)")
	NewOrderCmd.Flags().StringVar(&optionOrderSymbol, "symbol", "", "Order symbol")
	NewOrderCmd.Flags().Int64Var(&optionOrderQuantity, "quantity", 1, "Order quantity")
	NewOrderCmd.Flags().StringVar(&optionOrderID, "id", "", "Order id (uuid autogenerated if not given)")
	NewOrderCmd.Flags().StringVar(&optionOrderExpiry, "expiry", "day", "Order expiry (day, good_till_cancel ... etc)")
	NewOrderCmd.Flags().Float64Var(&optionOrderPrice, "price", 0.0, "Order price")
	NewOrderCmd.Flags().StringVar(&optionOrderOrigination, "origination", "", "Order origination")
	NewOrderCmd.Flags().StringSliceVar(&optionOrderAttributeType, "attribute-type", []string{}, "Order attribute types")
	NewOrderCmd.Flags().StringSliceVar(&optionOrderAttributeValue, "attribute-value", []string{}, "Order attribute value")

	NewOrderCmd.Flags().StringSliceVar(&optionPartyIDs, "party-id", []string{}, "Order party ids")
	NewOrderCmd.Flags().StringSliceVar(&optionPartyIDSources, "party-id-source", []string{}, "Order party id sources")
	NewOrderCmd.Flags().StringSliceVar(&optionPartyRoles, "party-role", []string{}, "Order party roles")
	NewOrderCmd.Flags().StringSliceVar(&optionPartyRoleQualifiers, "party-role-qualifier", []string{}, "Order party role qualifiers")
	NewOrderCmd.Flags().StringSliceVar(&optionPartySubIDs, "party-sub-ids", []string{}, "Order party sub ids (space separated)")
	NewOrderCmd.Flags().StringSliceVar(&optionPartySubIDTypes, "party-sub-id-types", []string{}, "Order party sub id types (space separated)")

	NewOrderCmd.Flags().IntVar(&optionExecReports, "exec-reports", 1, "Expect given number of execution reports before logging out (0 wait indefinitely)")
	NewOrderCmd.Flags().DurationVar(&optionExecReportsTimeout, "exec-reports-timeout", 5*time.Second, "Log out if execution reports not received within timeout (0s wait indefinitely)")
	NewOrderCmd.Flags().BoolVar(&optionExecReportsTimeoutReset, "exec-reports-timeout-reset", false, "Reset execution reports timeout each time an execution report is received")

	NewOrderCmd.MarkFlagRequired("side")
	NewOrderCmd.MarkFlagRequired("type")
	NewOrderCmd.MarkFlagRequired("symbol")
	NewOrderCmd.MarkFlagRequired("quantity")

	NewOrderCmd.RegisterFlagCompletionFunc("side", complete.OrderSide)
	NewOrderCmd.RegisterFlagCompletionFunc("type", complete.OrderType)
	NewOrderCmd.RegisterFlagCompletionFunc("expiry", complete.OrderTimeInForce)
	NewOrderCmd.RegisterFlagCompletionFunc("symbol", cobra.NoFileCompletions)
	NewOrderCmd.RegisterFlagCompletionFunc("origination", complete.OrderOriginationRole)
	NewOrderCmd.RegisterFlagCompletionFunc("attribute-type", complete.OrderAttributeType)
	NewOrderCmd.RegisterFlagCompletionFunc("attribute-value", cobra.NoFileCompletions)
	NewOrderCmd.RegisterFlagCompletionFunc("party-id", cobra.NoFileCompletions)
	NewOrderCmd.RegisterFlagCompletionFunc("party-id-source", complete.OrderPartyIDSource)
	NewOrderCmd.RegisterFlagCompletionFunc("party-sub-ids", cobra.NoFileCompletions)
	NewOrderCmd.RegisterFlagCompletionFunc("party-sub-id-types", complete.OrderPartySubIDTypes)
	NewOrderCmd.RegisterFlagCompletionFunc("party-role", complete.OrderPartyIDRole)
	NewOrderCmd.RegisterFlagCompletionFunc("party-role-qualifier", complete.OrderPartyRoleQualifier)
}

func Validate(cmd *cobra.Command, args []string) error {
	sides := utils.PrettyOptionValues(dict.OrderSides)
	search := utils.Search(sides, strings.ToLower(optionOrderSide))
	if search < 0 {
		return errors.OptionOrderSideUnknown
	}

	types := utils.PrettyOptionValues(dict.OrderTypes)
	search = utils.Search(types, strings.ToLower(optionOrderType))
	if search < 0 {
		return errors.OptionOrderTypeUnknown
	}

	if len(optionOrderID) == 0 {
		uid := uuid.New()
		optionOrderID = uid.String()
	}

	if len(optionOrderOrigination) > 0 {
		originations := utils.PrettyOptionValues(dict.OrderOriginations)
		search = utils.Search(originations, strings.ToLower(optionOrderOrigination))
		if search < 0 {
			return errors.OptionOrderOriginationUnknown
		}
	}

	if strings.ToLower(optionOrderType) == "market" && optionOrderPrice > 0 {
		return errors.OptionsInvalidMarketPrice
	} else if strings.ToLower(optionOrderType) != "market" && optionOrderPrice == 0 {
		return errors.OptionsNoPriceGiven
	}

	// Attributes
	if len(optionOrderAttributeType) != len(optionOrderAttributeValue) &&
		len(optionOrderAttributeValue) > 0 {
		return fmt.Errorf("%v: you must provide the same number of --attribute-type and --attribute-values", errors.OptionsInconsistentValues)
	}

	attributeTypes := utils.PrettyOptionValues(dict.OrderAttributeTypes)
	for k := range optionOrderAttributeType {
		search = utils.Search(attributeTypes, strings.ToLower(optionOrderAttributeType[k]))
		if search < 0 {
			return fmt.Errorf("%w: `%s`", errors.OptionOrderAttributeTypeUnkonwn, optionOrderAttributeType[k])
		}
	}

	// Parties
	if len(optionPartyIDs) != len(optionPartyIDSources) ||
		len(optionPartyIDs) != len(optionPartyRoles) ||
		len(optionPartyIDSources) != len(optionPartyRoles) {
		return fmt.Errorf("%v: you must provide the same number of --party-id, --party-id-source, --party-sub-id and --party-role", errors.OptionsInconsistentValues)
	}

	if len(optionPartyRoleQualifiers) > 0 && len(optionPartyRoleQualifiers) != len(optionPartyIDs) {
		return fmt.Errorf("%v: you must provide the same number of --party-id and --party-role-qualifier (%d != %d), %#v", errors.OptionsInconsistentValues, len(optionPartyIDs), len(optionPartyRoleQualifiers), optionPartyRoleQualifiers)
	}

	if len(optionPartySubIDs) > 0 && len(optionPartySubIDs) != len(optionPartyIDs) {
		return fmt.Errorf("%v: you must provide the same number of --party-id and --party-sub-ids", errors.OptionsInconsistentValues)
	}

	if len(optionPartySubIDTypes) > 0 && len(optionPartySubIDTypes) != len(optionPartyIDs) {
		return fmt.Errorf("%v: you must provide the same number of --party-id and --party-sub-id-types", errors.OptionsInconsistentValues)
	}

	sources := utils.PrettyOptionValues(dict.PartyIDSources)
	for k := range optionPartyIDSources {
		search = utils.Search(sources, strings.ToLower(optionPartyIDSources[k]))
		if search < 0 {
			return fmt.Errorf("%w: `%s`", errors.OptionOrderIDSourceUnknown, optionPartyIDSources[k])
		}
	}

	roles := utils.PrettyOptionValues(dict.PartyRoles)
	for k := range optionPartyRoles {
		search = utils.Search(roles, strings.ToLower(optionPartyRoles[k]))
		if search < 0 {
			return fmt.Errorf("%w: `%s`", errors.OptionOrderRoleUnknown, optionPartyRoles[k])
		}
	}

	roleQualifiers := utils.PrettyOptionValues(dict.PartyRoleQualifiers)
	for k := range optionPartyRoleQualifiers {
		search = utils.Search(roleQualifiers, strings.ToLower(optionPartyRoleQualifiers[k]))
		if search < 0 {
			return fmt.Errorf("%w: `%s`", errors.OptionOrderRoleQualifierUnknown, optionPartyRoleQualifiers[k])
		}
	}

	// Sub Parties
	partySubIDTypes := utils.PrettyOptionValues(dict.PartySubIDTypes)
	for k := range optionPartySubIDs {
		var subIDs, subIDTypes []string

		if len(optionPartySubIDs) > 0 {
			subIDs = strings.Split(optionPartySubIDs[k], " ")
		}

		if len(optionPartySubIDTypes) > 0 {
			subIDTypes = strings.Split(optionPartySubIDTypes[k], " ")
		}

		if len(subIDs) > 0 && len(subIDTypes) > 0 && len(subIDs) != len(subIDTypes) {
			return fmt.Errorf("%v: %s occurence of --party-sub-ids and --party-sub-id-types do not have same number of elements (space separated)", errors.OptionsInconsistentValues, humanize.Ordinal(k))
		}

		if len(subIDTypes) > 0 {
			for kk := range subIDTypes {
				search = utils.Search(partySubIDTypes, strings.ToLower(subIDTypes[kk]))
				if search < 0 {
					return fmt.Errorf("%w: `%s`", errors.OptionPartySubIDTypeUnknown, subIDTypes[kk])
				}
			}
		}
	}

	return nil
}

func Execute(cmd *cobra.Command, args []string) error {
	options := config.GetOptions()
	logger := config.GetLogger()

	context, err := config.GetCurrentContext()
	if err != nil {
		return err
	}

	sessions, err := context.GetSessions()
	if err != nil {
		return err
	}

	session := sessions[0]
	initiatior, err := context.GetInitiator()
	if err != nil {
		return err
	}

	transportDict, appDict, err := session.GetFIXDictionaries()
	if err != nil {
		return err
	}

	settings, err := context.ToQuickFixInitiatorSettings()
	if err != nil {
		return err
	}

	app := application.NewNewOrder()
	app.Logger = logger
	app.Settings = settings
	app.TransportDataDictionary = transportDict
	app.AppDataDictionary = appDict

	var quickfixLogger *zerolog.Logger
	if options.QuickFixLogging {
		quickfixLogger = logger
	}

	// Choose right timeout cli option > config > default value (5s)
	var timeout time.Duration
	if options.Timeout != time.Duration(0) {
		timeout = options.Timeout
	} else if initiatior.SocketTimeout != time.Duration(0) {
		timeout = initiatior.SocketTimeout
	} else {
		timeout = 5 * time.Second
	}

	init, err := initiator.Initiate(app, settings, quickfixLogger)
	if err != nil {
		return err
	}

	// Start session
	err = init.Start()
	if err != nil {
		return err
	}

	// Defer stopping initiator
	defer init.Stop()

	// Wait for session connection
	select {
	case <-time.After(timeout):
		return errors.ConnectionTimeout
	case _, ok := <-app.Connected:
		if !ok {
			return errors.FixLogout
		}
	}

	// Prepare order
	order, err := buildMessage(*session)
	if err != nil {
		return err
	}

	// Send the order
	err = quickfix.Send(order)
	if err != nil {
		return err
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	execReports := 0
	var waitTimeout <-chan time.Time
	if optionExecReportsTimeout > 0 {
		waitTimeout = time.After(optionExecReportsTimeout)
	} else {
		waitTimeout = make(<-chan time.Time)
	}

LOOP:
	for {
		select {
		case signal := <-interrupt:
			logger.Debug().Msgf("Received signal: %s", signal)
			break LOOP

		case <-waitTimeout:
			logger.Warn().Msgf("Timeout while expecting execution reports (%d/%d)", execReports, optionExecReports)
			break LOOP

		case msg, ok := <-app.FromAppChan:
			if !ok {
				break LOOP
			}

			if err := processReponse(app, msg); err != nil {
				if errors.Is(err, quickfix.InvalidMessageType()) {
					continue LOOP
				}

				return err
			}

			// Reset timeout
			if optionExecReportsTimeoutReset && optionExecReportsTimeout > 0 {
				waitTimeout = time.After(optionExecReportsTimeout)
			}

			execReports = execReports + 1
		}

		if optionExecReports != 0 && execReports >= optionExecReports {
			break LOOP
		}
	}

	return nil
}

func buildMessage(session config.Session) (quickfix.Messagable, error) {
	eside, err := dict.OrderSideStringToEnum(optionOrderSide)
	if err != nil {
		return nil, err
	}

	etype, err := dict.OrderTypeStringToEnum(optionOrderType)
	if err != nil {
		return nil, err
	}

	eExpiry, err := dict.OrderTimeInForceStringToEnum(optionOrderExpiry)
	if err != nil {
		return nil, err
	}

	// Prepare order
	clordid := field.NewClOrdID(optionOrderID)
	ordtype := field.NewOrdType(etype)
	transactime := field.NewTransactTime(time.Now())
	ordside := field.NewSide(eside)

	// Message
	message := quickfix.NewMessage()
	header := fixt11.NewHeader(&message.Header)

	switch session.BeginString {
	case quickfix.BeginStringFIXT11:
		switch session.DefaultApplVerID {
		case "FIX.5.0SP2":
			header.Set(field.NewMsgType(enum.MsgType_ORDER_SINGLE))
			message.Body.Set(clordid)
			message.Body.Set(ordside)
			message.Body.Set(transactime)
			message.Body.Set(ordtype)

			// Parties
			NewNoPartySubIDsRepeatingGroup := func() *quickfix.RepeatingGroup {
				return quickfix.NewRepeatingGroup(
					tag.NoPartySubIDs,
					quickfix.GroupTemplate{
						quickfix.GroupElement(tag.PartySubID),
						quickfix.GroupElement(tag.PartySubIDType),
					},
				)
			}
			parties := quickfix.NewRepeatingGroup(
				tag.NoPartyIDs,
				quickfix.GroupTemplate{
					quickfix.GroupElement(tag.PartyID),
					quickfix.GroupElement(tag.PartyIDSource),
					quickfix.GroupElement(tag.PartyRole),
					NewNoPartySubIDsRepeatingGroup(),
				},
			)

			for i := range optionPartyIDs {
				party := parties.Add()

				party.Set(field.NewPartyID(optionPartyIDs[i]))
				party.Set(field.NewPartyIDSource(enum.PartyIDSource(dict.PartyIDSources[strings.ToUpper(optionPartyIDSources[i])])))
				party.Set(field.NewPartyRole(enum.PartyRole(dict.PartyRoles[strings.ToUpper(optionPartyRoles[i])])))

				// Role Qualifier
				if len(optionPartyRoleQualifiers) > 0 && len(optionPartyRoleQualifiers[i]) > 0 {
					party.Set(field.NewPartyRoleQualifier(utils.Must(strconv.Atoi(string(dict.PartyRoleQualifiers[strings.ToUpper(optionPartyRoleQualifiers[i])])))))
				}

				if len(optionPartySubIDs) == len(optionPartyIDs) ||
					len(optionPartySubIDTypes) == len(optionPartyIDs) {
					var partySubIDs, partySubIDTypes []string

					if len(optionPartySubIDs) > 0 {
						partySubIDs = strings.Split(optionPartySubIDs[i], " ")
					}

					if len(optionPartySubIDTypes) > 0 {
						partySubIDTypes = strings.Split(optionPartySubIDTypes[i], " ")
					}

					if len(partySubIDs) > 0 || len(partySubIDTypes) > 0 {
						subIDs := NewNoPartySubIDsRepeatingGroup()
						for k := range partySubIDs {
							subID := subIDs.Add()
							mustAdd := false
							if len(partySubIDs) > 0 {
								mustAdd = true
								subID.Set(field.NewPartySubID(partySubIDs[k]))
							}
							if len(partySubIDTypes) > 0 {
								mustAdd = true
								subID.Set(field.NewPartySubIDType(enum.PartySubIDType(dict.PartySubIDTypes[strings.ToUpper(partySubIDTypes[k])])))
							}
							if mustAdd {
								party.SetGroup(subIDs)
							}
						}
					}

				}
			}
			message.Body.SetGroup(parties)

			// Attributes
			attributes := quickfix.NewRepeatingGroup(
				tag.NoOrderAttributes,
				quickfix.GroupTemplate{
					quickfix.GroupElement(tag.OrderAttributeType),
					quickfix.GroupElement(tag.OrderAttributeValue),
				},
			)

			for i := range optionOrderAttributeType {
				attribute := attributes.Add()
				attribute.Set(field.NewOrderAttributeType(enum.OrderAttributeType(dict.OrderAttributeTypes[strings.ToUpper(optionOrderAttributeType[i])])))

				if len(optionOrderAttributeValue) > 0 {
					if len(optionOrderAttributeValue[i]) > 0 {
						attribute.SetString(tag.OrderAttributeValue, optionOrderAttributeValue[i])
					}
				}
			}

			message.Body.SetGroup(attributes)

		default:
			return nil, errors.FixVersionNotImplemented
		}
	default:
		return nil, errors.FixVersionNotImplemented
	}

	utils.QuickFixMessagePartSetString(&message.Header, session.TargetCompID, field.NewTargetCompID)
	utils.QuickFixMessagePartSetString(&message.Header, session.TargetSubID, field.NewTargetSubID)
	utils.QuickFixMessagePartSetString(&message.Header, session.SenderCompID, field.NewSenderCompID)
	utils.QuickFixMessagePartSetString(&message.Header, session.SenderSubID, field.NewSenderSubID)

	message.Body.Set(field.NewHandlInst(enum.HandlInst_AUTOMATED_EXECUTION_ORDER_PRIVATE_NO_BROKER_INTERVENTION))
	message.Body.Set(field.NewSymbol(optionOrderSymbol))
	message.Body.Set(field.NewOrderQty(decimal.NewFromInt(optionOrderQuantity), 2))
	message.Body.Set(field.NewTimeInForce(eExpiry))

	if etype != enum.OrdType_MARKET {
		message.Body.Set(field.NewPrice(decimal.NewFromFloat(optionOrderPrice), 2))
	}

	if len(optionOrderOrigination) > 0 {
		message.Body.Set(field.NewOrderOrigination(enum.OrderOrigination(dict.OrderOriginations[strings.ToUpper(optionOrderOrigination)])))
	}

	return message, nil
}

func processReponse(app *application.NewOrder, msg *quickfix.Message) error {
	msgType := field.MsgTypeField{}
	ordStatus := field.OrdStatusField{}
	text := field.TextField{}

	// MsgType
	err := msg.Header.GetField(tag.MsgType, &msgType)
	if err != nil {
		return err
	} else if msgType.Value() != enum.MsgType_EXECUTION_REPORT {
		return quickfix.InvalidMessageType()
	}

	if msgType.Value() == enum.MsgType_REJECT {
		return fmt.Errorf("%w: %s", errors.FixOrderRejected, text.String())
	}

	// OrdStatus
	msg.Body.GetField(tag.OrdStatus, &ordStatus)
	if err != nil {
		return err
	}

	switch ordStatus.Value() {
	case enum.OrdStatus_NEW:
		fallthrough
	case enum.OrdStatus_PARTIALLY_FILLED:
		fallthrough
	case enum.OrdStatus_FILLED:
		fallthrough
	case enum.OrdStatus_DONE_FOR_DAY:
		fallthrough
	case enum.OrdStatus_CANCELED:
		fallthrough
	case enum.OrdStatus_REPLACED:
		fallthrough
	case enum.OrdStatus_PENDING_CANCEL:
		fallthrough
	case enum.OrdStatus_STOPPED:
		fallthrough
	case enum.OrdStatus_SUSPENDED:
		fallthrough
	case enum.OrdStatus_PENDING_NEW:
		fallthrough
	case enum.OrdStatus_CALCULATED:
		fallthrough
	case enum.OrdStatus_EXPIRED:
		fallthrough
	case enum.OrdStatus_ACCEPTED_FOR_BIDDING:
		fallthrough
	case enum.OrdStatus_PENDING_REPLACE:
		app.WriteMessageBodyAsTable(os.Stdout, msg)

	case enum.OrdStatus_REJECTED:
		msg.Body.GetField(tag.Text, &text)
		if err != nil {
			return errors.FixOrderRejected
		}

		return fmt.Errorf("%w: %s", errors.FixOrderRejected, text.String())

	default:
		if len(text.String()) > 0 {
			return fmt.Errorf("%w: %s", errors.FixOrderStatusUnknown, text.String())
		}
	}

	return nil
}

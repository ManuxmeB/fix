package dict

import (
	"github.com/quickfixgo/enum"
)

var PartyIDSources = map[string]enum.PartyIDSource{
	"KOREAN_INVESTOR_ID": enum.PartyIDSource_KOREAN_INVESTOR_ID,
	"TAIWANESE_QUALIFIED_FOREIGN_INVESTOR_ID_QFII_FID": enum.PartyIDSource_TAIWANESE_QUALIFIED_FOREIGN_INVESTOR_ID_QFII_FID,
	"TAIWANESE_TRADING_ACCT":                           enum.PartyIDSource_TAIWANESE_TRADING_ACCT,
	"MALAYSIAN_CENTRAL_DEPOSITORY":                     enum.PartyIDSource_MALAYSIAN_CENTRAL_DEPOSITORY,
	"CHINESE_INVESTOR_ID":                              enum.PartyIDSource_CHINESE_INVESTOR_ID,
	"UK_NATIONAL_INSURANCE_OR_PENSION_NUMBER":          enum.PartyIDSource_UK_NATIONAL_INSURANCE_OR_PENSION_NUMBER,
	"US_SOCIAL_SECURITY_NUMBER":                        enum.PartyIDSource_US_SOCIAL_SECURITY_NUMBER,
	"US_EMPLOYER_OR_TAX_ID_NUMBER":                     enum.PartyIDSource_US_EMPLOYER_OR_TAX_ID_NUMBER,
	"AUSTRALIAN_BUSINESS_NUMBER":                       enum.PartyIDSource_AUSTRALIAN_BUSINESS_NUMBER,
	"AUSTRALIAN_TAX_FILE_NUMBER":                       enum.PartyIDSource_AUSTRALIAN_TAX_FILE_NUMBER,
	"BIC":                                              enum.PartyIDSource_BIC,
	"GENERALLY_ACCEPTED_MARKET_PARTICIPANT_IDENTIFIER": enum.PartyIDSource_GENERALLY_ACCEPTED_MARKET_PARTICIPANT_IDENTIFIER,
	"PROPRIETARY":                                      enum.PartyIDSource_PROPRIETARY,
	"ISO_COUNTRY_CODE":                                 enum.PartyIDSource_ISO_COUNTRY_CODE,
	"SETTLEMENT_ENTITY_LOCATION":                       enum.PartyIDSource_SETTLEMENT_ENTITY_LOCATION,
	"MARKET_IDENTIFIER_CODE":                           enum.PartyIDSource_MARKET_IDENTIFIER_CODE,
	"CSD_PARTICIPANT_MEMBER_CODE":                      enum.PartyIDSource_CSD_PARTICIPANT_MEMBER_CODE,
	"DIRECTED_BROKER_THREE_CHARACTER_ACRONYM_AS_DEFINED_IN_ISITC_ETC_BEST_PRACTICE_GUIDELINES_DOCUMENT": enum.PartyIDSource_DIRECTED_BROKER_THREE_CHARACTER_ACRONYM_AS_DEFINED_IN_ISITC_ETC_BEST_PRACTICE_GUIDELINES_DOCUMENT,
	"TAX_ID":                               enum.PartyIDSource_TAX_ID,
	"AUSTRALIAN_COMPANY_NUMBER":            enum.PartyIDSource_AUSTRALIAN_COMPANY_NUMBER,
	"AUSTRALIAN_REGISTERED_BODY_NUMBER":    enum.PartyIDSource_AUSTRALIAN_REGISTERED_BODY_NUMBER,
	"CFTC_REPORTING_FIRM_IDENTIFIER":       enum.PartyIDSource_CFTC_REPORTING_FIRM_IDENTIFIER,
	"LEGAL_ENTITY_IDENTIFIER":              enum.PartyIDSource_LEGAL_ENTITY_IDENTIFIER,
	"INTERIM_IDENTIFIER":                   enum.PartyIDSource_INTERIM_IDENTIFIER,
	"SHORT_CODE_IDENTIFIER":                enum.PartyIDSource_SHORT_CODE_IDENTIFIER,
	"NATIONAL_ID_OF_NATURAL_PERSON":        enum.PartyIDSource_NATIONAL_ID_OF_NATURAL_PERSON,
	"INDIA_PERMANENT_ACCOUNT_NUMBER":       enum.PartyIDSource_INDIA_PERMANENT_ACCOUNT_NUMBER,
	"FIRM_DESIGNATED_IDENTIFIER":           enum.PartyIDSource_FIRM_DESIGNATED_IDENTIFIER,
	"SPECIAL_SEGREGATED_ACCOUNT_ID":        enum.PartyIDSource_SPECIAL_SEGREGATED_ACCOUNT_ID,
	"MASTER_SPECIAL_SEGREGATED_ACCOUNT_ID": enum.PartyIDSource_MASTER_SPECIAL_SEGREGATED_ACCOUNT_ID,
}

var PartyRoles = map[string]enum.PartyRole{
	"EXECUTING_FIRM":                              enum.PartyRole_EXECUTING_FIRM,
	"SETTLEMENT_LOCATION":                         enum.PartyRole_SETTLEMENT_LOCATION,
	"MARGIN_ACCOUNT":                              enum.PartyRole_MARGIN_ACCOUNT,
	"COLLATERAL_ASSET_ACCOUNT":                    enum.PartyRole_COLLATERAL_ASSET_ACCOUNT,
	"DATA_REPOSITORY":                             enum.PartyRole_DATA_REPOSITORY,
	"CALCULATION_AGENT":                           enum.PartyRole_CALCULATION_AGENT,
	"SENDER_OF_EXERCISE_NOTICE":                   enum.PartyRole_SENDER_OF_EXERCISE_NOTICE,
	"RECEIVER_OF_EXERCISE_NOTICE":                 enum.PartyRole_RECEIVER_OF_EXERCISE_NOTICE,
	"RATE_REFERENCE_BANK":                         enum.PartyRole_RATE_REFERENCE_BANK,
	"CORRESPONDENT":                               enum.PartyRole_CORRESPONDENT,
	"BENEFICIARYS_BANK_OR_DEPOSITORY_INSTITUTION": enum.PartyRole_BENEFICIARYS_BANK_OR_DEPOSITORY_INSTITUTION,
	"ORDER_ORIGINATION_TRADER":                    enum.PartyRole_ORDER_ORIGINATION_TRADER,
	"BORROWER":                                    enum.PartyRole_BORROWER,
	"PRIMARY_OBLIGATOR":                           enum.PartyRole_PRIMARY_OBLIGATOR,
	"GUARANTOR":                                   enum.PartyRole_GUARANTOR,
	"EXCLUDED_REFERENCE_ENTITY":                   enum.PartyRole_EXCLUDED_REFERENCE_ENTITY,
	"DETERMINING_PARTY":                           enum.PartyRole_DETERMINING_PARTY,
	"HEDGING_PARTY":                               enum.PartyRole_HEDGING_PARTY,
	"REPORTING_ENTITY":                            enum.PartyRole_REPORTING_ENTITY,
	"SALES_PERSON":                                enum.PartyRole_SALES_PERSON,
	"OPERATOR":                                    enum.PartyRole_OPERATOR,
	"CENTRAL_SECURITIES_DEPOSITORY_119":           enum.PartyRole_CENTRAL_SECURITIES_DEPOSITORY_119,
	"EXECUTING_TRADER":                            enum.PartyRole_EXECUTING_TRADER,
	"INTERNATIONAL_CENTRAL_SECURITIES_DEPOSITORY": enum.PartyRole_INTERNATIONAL_CENTRAL_SECURITIES_DEPOSITORY,
	"TRADING_SUB_ACCOUNT":                         enum.PartyRole_TRADING_SUB_ACCOUNT,
	"INVESTMENT_DECISION_MAKER":                   enum.PartyRole_INVESTMENT_DECISION_MAKER,
	"PUBLISHING_INTERMEDIARY":                     enum.PartyRole_PUBLISHING_INTERMEDIARY,
	"CENTRAL_SECURITIES_DEPOSITORY_124":           enum.PartyRole_CENTRAL_SECURITIES_DEPOSITORY_124,
	"ISSUER":                                      enum.PartyRole_ISSUER,
	"CONTRA_CUSTOMER_ACCOUNT":                     enum.PartyRole_CONTRA_CUSTOMER_ACCOUNT,
	"CONTRA_INVESTMENT_DECISION_MAKER":            enum.PartyRole_CONTRA_INVESTMENT_DECISION_MAKER,
	"ORDER_ORIGINATION_FIRM":                      enum.PartyRole_ORDER_ORIGINATION_FIRM,
	"GIVEUP_CLEARING_FIRM":                        enum.PartyRole_GIVEUP_CLEARING_FIRM,
	"CORRESPONDANT_CLEARING_FIRM":                 enum.PartyRole_CORRESPONDANT_CLEARING_FIRM,
	"EXECUTING_SYSTEM":                            enum.PartyRole_EXECUTING_SYSTEM,
	"CONTRA_FIRM":                                 enum.PartyRole_CONTRA_FIRM,
	"CONTRA_CLEARING_FIRM":                        enum.PartyRole_CONTRA_CLEARING_FIRM,
	"SPONSORING_FIRM":                             enum.PartyRole_SPONSORING_FIRM,
	"BROKER_OF_CREDIT":                            enum.PartyRole_BROKER_OF_CREDIT,
	"UNDERLYING_CONTRA_FIRM":                      enum.PartyRole_UNDERLYING_CONTRA_FIRM,
	"CLEARING_ORGANIZATION":                       enum.PartyRole_CLEARING_ORGANIZATION,
	"EXCHANGE":                                    enum.PartyRole_EXCHANGE,
	"CUSTOMER_ACCOUNT":                            enum.PartyRole_CUSTOMER_ACCOUNT,
	"CORRESPONDENT_CLEARING_ORGANIZATION":         enum.PartyRole_CORRESPONDENT_CLEARING_ORGANIZATION,
	"CORRESPONDENT_BROKER":                        enum.PartyRole_CORRESPONDENT_BROKER,
	"BUYER_SELLER":                                enum.PartyRole_BUYER_SELLER,
	"CUSTODIAN":                                   enum.PartyRole_CUSTODIAN,
	"INTERMEDIARY":                                enum.PartyRole_INTERMEDIARY,
	"CLIENT_ID":                                   enum.PartyRole_CLIENT_ID,
	"AGENT":                                       enum.PartyRole_AGENT,
	"SUB_CUSTODIAN":                               enum.PartyRole_SUB_CUSTODIAN,
	"BENEFICIARY":                                 enum.PartyRole_BENEFICIARY,
	"INTERESTED_PARTY":                            enum.PartyRole_INTERESTED_PARTY,
	"REGULATORY_BODY":                             enum.PartyRole_REGULATORY_BODY,
	"LIQUIDITY_PROVIDER":                          enum.PartyRole_LIQUIDITY_PROVIDER,
	"ENTERING_TRADER":                             enum.PartyRole_ENTERING_TRADER,
	"CONTRA_TRADER":                               enum.PartyRole_CONTRA_TRADER,
	"POSITION_ACCOUNT":                            enum.PartyRole_POSITION_ACCOUNT,
	"CONTRA_INVESTOR_ID":                          enum.PartyRole_CONTRA_INVESTOR_ID,
	"CLEARING_FIRM":                               enum.PartyRole_CLEARING_FIRM,
	"TRANSFER_TO_FIRM":                            enum.PartyRole_TRANSFER_TO_FIRM,
	"CONTRA_POSITION_ACCOUNT":                     enum.PartyRole_CONTRA_POSITION_ACCOUNT,
	"CONTRA_EXCHANGE":                             enum.PartyRole_CONTRA_EXCHANGE,
	"INTERNAL_CARRY_ACCOUNT":                      enum.PartyRole_INTERNAL_CARRY_ACCOUNT,
	"ORDER_ENTRY_OPERATOR_ID":                     enum.PartyRole_ORDER_ENTRY_OPERATOR_ID,
	"SECONDARY_ACCOUNT_NUMBER":                    enum.PartyRole_SECONDARY_ACCOUNT_NUMBER,
	"FOREIGN_FIRM":                                enum.PartyRole_FOREIGN_FIRM,
	"THIRD_PARTY_ALLOCATION_FIRM":                 enum.PartyRole_THIRD_PARTY_ALLOCATION_FIRM,
	"CLAIMING_ACCOUNT":                            enum.PartyRole_CLAIMING_ACCOUNT,
	"ASSET_MANAGER":                               enum.PartyRole_ASSET_MANAGER,
	"INVESTOR_ID":                                 enum.PartyRole_INVESTOR_ID,
	"PLEDGOR_ACCOUNT":                             enum.PartyRole_PLEDGOR_ACCOUNT,
	"PLEDGEE_ACCOUNT":                             enum.PartyRole_PLEDGEE_ACCOUNT,
	"LARGE_TRADER_REPORTABLE_ACCOUNT":             enum.PartyRole_LARGE_TRADER_REPORTABLE_ACCOUNT,
	"TRADER_MNEMONIC":                             enum.PartyRole_TRADER_MNEMONIC,
	"SENDER_LOCATION":                             enum.PartyRole_SENDER_LOCATION,
	"SESSION_ID":                                  enum.PartyRole_SESSION_ID,
	"ACCEPTABLE_COUNTERPARTY":                     enum.PartyRole_ACCEPTABLE_COUNTERPARTY,
	"UNACCEPTABLE_COUNTERPARTY":                   enum.PartyRole_UNACCEPTABLE_COUNTERPARTY,
	"ENTERING_UNIT":                               enum.PartyRole_ENTERING_UNIT,
	"EXECUTING_UNIT":                              enum.PartyRole_EXECUTING_UNIT,
	"INTRODUCING_FIRM":                            enum.PartyRole_INTRODUCING_FIRM,
	"INTRODUCING_BROKER":                          enum.PartyRole_INTRODUCING_BROKER,
	"QUOTE_ORIGINATOR":                            enum.PartyRole_QUOTE_ORIGINATOR,
	"REPORT_ORIGINATOR":                           enum.PartyRole_REPORT_ORIGINATOR,
	"SYSTEMATIC_INTERNALISER":                     enum.PartyRole_SYSTEMATIC_INTERNALISER,
	"MULTILATERAL_TRADING_FACILITY":               enum.PartyRole_MULTILATERAL_TRADING_FACILITY,
	"REGULATED_MARKET":                            enum.PartyRole_REGULATED_MARKET,
	"MARKET_MAKER":                                enum.PartyRole_MARKET_MAKER,
	"INVESTMENT_FIRM":                             enum.PartyRole_INVESTMENT_FIRM,
	"HOST_COMPETENT_AUTHORITY":                    enum.PartyRole_HOST_COMPETENT_AUTHORITY,
	"HOME_COMPETENT_AUTHORITY":                    enum.PartyRole_HOME_COMPETENT_AUTHORITY,
	"ENTERING_FIRM":                               enum.PartyRole_ENTERING_FIRM,
	"COMPETENT_AUTHORITY_OF_THE_MOST_RELEVANT_MARKET_IN_TERMS_OF_LIQUIDITY": enum.PartyRole_COMPETENT_AUTHORITY_OF_THE_MOST_RELEVANT_MARKET_IN_TERMS_OF_LIQUIDITY,
	"COMPETENT_AUTHORITY_OF_THE_TRANSACTION":                                enum.PartyRole_COMPETENT_AUTHORITY_OF_THE_TRANSACTION,
	"REPORTING_INTERMEDIARY":                                                enum.PartyRole_REPORTING_INTERMEDIARY,
	"EXECUTION_VENUE":                                                       enum.PartyRole_EXECUTION_VENUE,
	"MARKET_DATA_ENTRY_ORIGINATOR":                                          enum.PartyRole_MARKET_DATA_ENTRY_ORIGINATOR,
	"LOCATION_ID":                                                           enum.PartyRole_LOCATION_ID,
	"DESK_ID":                                                               enum.PartyRole_DESK_ID,
	"MARKET_DATA_MARKET":                                                    enum.PartyRole_MARKET_DATA_MARKET,
	"ALLOCATION_ENTITY":                                                     enum.PartyRole_ALLOCATION_ENTITY,
	"PRIME_BROKER_PROVIDING_GENERAL_TRADE_SERVICES":                         enum.PartyRole_PRIME_BROKER_PROVIDING_GENERAL_TRADE_SERVICES,
	"LOCATE":                             enum.PartyRole_LOCATE,
	"STEP_OUT_FIRM":                      enum.PartyRole_STEP_OUT_FIRM,
	"BROKER_CIENT_ID":                    enum.PartyRole_BROKER_CIENT_ID,
	"CENTRAL_REGISTRATION_DEPOSITORY":    enum.PartyRole_CENTRAL_REGISTRATION_DEPOSITORY,
	"CLEARING_ACCOUNT":                   enum.PartyRole_CLEARING_ACCOUNT,
	"ACCEPTABLE_SETTLING_COUNTERPARTY":   enum.PartyRole_ACCEPTABLE_SETTLING_COUNTERPARTY,
	"UNACCEPTABLE_SETTLING_COUNTERPARTY": enum.PartyRole_UNACCEPTABLE_SETTLING_COUNTERPARTY,
	"CLS_MEMBER_BANK":                    enum.PartyRole_CLS_MEMBER_BANK,
	"IN_CONCERT_GROUP":                   enum.PartyRole_IN_CONCERT_GROUP,
	"IN_CONCERT_CONTROLLING_ENTITY":      enum.PartyRole_IN_CONCERT_CONTROLLING_ENTITY,
	"LARGE_POSITIONS_REPORTING_ACCOUNT":  enum.PartyRole_LARGE_POSITIONS_REPORTING_ACCOUNT,
	"FUND_MANAGER_CLIENT_ID":             enum.PartyRole_FUND_MANAGER_CLIENT_ID,
	"SETTLEMENT_FIRM":                    enum.PartyRole_SETTLEMENT_FIRM,
	"SETTLEMENT_ACCOUNT":                 enum.PartyRole_SETTLEMENT_ACCOUNT,
	"REPORTING_MARKET_CENTER":            enum.PartyRole_REPORTING_MARKET_CENTER,
	"RELATED_REPORTING_MARKET_CENTER":    enum.PartyRole_RELATED_REPORTING_MARKET_CENTER,
	"AWAY_MARKET":                        enum.PartyRole_AWAY_MARKET,
	"GIVE_UP":                            enum.PartyRole_GIVE_UP,
	"TAKE_UP":                            enum.PartyRole_TAKE_UP,
	"GIVE_UP_CLEARING_FIRM":              enum.PartyRole_GIVE_UP_CLEARING_FIRM,
	"TAKE_UP_CLEARING_FIRM":              enum.PartyRole_TAKE_UP_CLEARING_FIRM,
	"ORIGINATING_MARKET":                 enum.PartyRole_ORIGINATING_MARKET,
}

var PartyRoleQualifiers = map[string]enum.PartyRoleQualifier{
	"AGENCY":                    enum.PartyRoleQualifier_AGENCY,
	"PRINCIPAL":                 enum.PartyRoleQualifier_PRINCIPAL,
	"ORIGINAL_TRADE_REPOSITORY": enum.PartyRoleQualifier_ORIGINAL_TRADE_REPOSITORY,
	"ADDITIONAL_INTERNATIONAL_TRADE_REPOSITORY": enum.PartyRoleQualifier_ADDITIONAL_INTERNATIONAL_TRADE_REPOSITORY,
	"ADDITIONAL_DOMESTIC_TRADE_REPOSITORY":      enum.PartyRoleQualifier_ADDITIONAL_DOMESTIC_TRADE_REPOSITORY,
	"RELATED_EXCHANGE":                          enum.PartyRoleQualifier_RELATED_EXCHANGE,
	"OPTIONS_EXCHANGE":                          enum.PartyRoleQualifier_OPTIONS_EXCHANGE,
	"SPECIFIED_EXCHANGE":                        enum.PartyRoleQualifier_SPECIFIED_EXCHANGE,
	"CONSTITUENT_EXCHANGE":                      enum.PartyRoleQualifier_CONSTITUENT_EXCHANGE,
	"EXEMPT_FROM_TRADE_REPORTING":               enum.PartyRoleQualifier_EXEMPT_FROM_TRADE_REPORTING,
	"CURRENT":                                   enum.PartyRoleQualifier_CURRENT,
	"NEW":                                       enum.PartyRoleQualifier_NEW,
	"RISKLESS_PRINCIPAL":                        enum.PartyRoleQualifier_RISKLESS_PRINCIPAL,
	"DESIGNATED_SPONSOR":                        enum.PartyRoleQualifier_DESIGNATED_SPONSOR,
	"SPECIALIST":                                enum.PartyRoleQualifier_SPECIALIST,
	"ALGORITHM":                                 enum.PartyRoleQualifier_ALGORITHM,
	"FIRM_OR_LEGAL_ENTITY":                      enum.PartyRoleQualifier_FIRM_OR_LEGAL_ENTITY,
	"NATURAL_PERSON":                            enum.PartyRoleQualifier_NATURAL_PERSON,
	"REGULAR_TRADER":                            enum.PartyRoleQualifier_REGULAR_TRADER,
	"HEAD_TRADER":                               enum.PartyRoleQualifier_HEAD_TRADER,
	"SUPERVISOR":                                enum.PartyRoleQualifier_SUPERVISOR,
	"TRI_PARTY":                                 enum.PartyRoleQualifier_TRI_PARTY,
	"LENDER":                                    enum.PartyRoleQualifier_LENDER,
	"GENERAL_CLEARING_MEMBER":                   enum.PartyRoleQualifier_GENERAL_CLEARING_MEMBER,
	"INDIVIDUAL_CLEARING_MEMBER":                enum.PartyRoleQualifier_INDIVIDUAL_CLEARING_MEMBER,
	"PREFERRED_MARKET_MAKER":                    enum.PartyRoleQualifier_PREFERRED_MARKET_MAKER,
	"DIRECTED_MARKET_MAKER":                     enum.PartyRoleQualifier_DIRECTED_MARKET_MAKER,
	"BANK":                                      enum.PartyRoleQualifier_BANK,
	"HUB":                                       enum.PartyRoleQualifier_HUB,
	"PRIMARY_TRADE_REPOSITORY":                  enum.PartyRoleQualifier_PRIMARY_TRADE_REPOSITORY,
}

var PartySubIDTypes = map[string]enum.PartySubIDType{
	"FIRM":                         enum.PartySubIDType_FIRM,
	"SECURITIES_ACCOUNT_NUMBER":    enum.PartySubIDType_SECURITIES_ACCOUNT_NUMBER,
	"REGISTRATION_NUMBER":          enum.PartySubIDType_REGISTRATION_NUMBER,
	"REGISTERED_ADDRESS_12":        enum.PartySubIDType_REGISTERED_ADDRESS_12,
	"REGULATORY_STATUS":            enum.PartySubIDType_REGULATORY_STATUS,
	"REGISTRATION_NAME":            enum.PartySubIDType_REGISTRATION_NAME,
	"CASH_ACCOUNT_NUMBER":          enum.PartySubIDType_CASH_ACCOUNT_NUMBER,
	"BIC":                          enum.PartySubIDType_BIC,
	"CSD_PARTICIPANT_MEMBER_CODE":  enum.PartySubIDType_CSD_PARTICIPANT_MEMBER_CODE,
	"REGISTERED_ADDRESS_18":        enum.PartySubIDType_REGISTERED_ADDRESS_18,
	"FUND_ACCOUNT_NAME":            enum.PartySubIDType_FUND_ACCOUNT_NAME,
	"PERSON":                       enum.PartySubIDType_PERSON,
	"TELEX_NUMBER":                 enum.PartySubIDType_TELEX_NUMBER,
	"FAX_NUMBER":                   enum.PartySubIDType_FAX_NUMBER,
	"SECURITIES_ACCOUNT_NAME":      enum.PartySubIDType_SECURITIES_ACCOUNT_NAME,
	"CASH_ACCOUNT_NAME":            enum.PartySubIDType_CASH_ACCOUNT_NAME,
	"DEPARTMENT":                   enum.PartySubIDType_DEPARTMENT,
	"LOCATION_DESK":                enum.PartySubIDType_LOCATION_DESK,
	"POSITION_ACCOUNT_TYPE":        enum.PartySubIDType_POSITION_ACCOUNT_TYPE,
	"SECURITY_LOCATE_ID":           enum.PartySubIDType_SECURITY_LOCATE_ID,
	"MARKET_MAKER":                 enum.PartySubIDType_MARKET_MAKER,
	"ELIGIBLE_COUNTERPARTY":        enum.PartySubIDType_ELIGIBLE_COUNTERPARTY,
	"SYSTEM":                       enum.PartySubIDType_SYSTEM,
	"PROFESSIONAL_CLIENT":          enum.PartySubIDType_PROFESSIONAL_CLIENT,
	"LOCATION":                     enum.PartySubIDType_LOCATION,
	"EXECUTION_VENUE":              enum.PartySubIDType_EXECUTION_VENUE,
	"CURRENCY_DELIVERY_IDENTIFIER": enum.PartySubIDType_CURRENCY_DELIVERY_IDENTIFIER,
	"APPLICATION":                  enum.PartySubIDType_APPLICATION,
	"FULL_LEGAL_NAME_OF_FIRM":      enum.PartySubIDType_FULL_LEGAL_NAME_OF_FIRM,
	"POSTAL_ADDRESS":               enum.PartySubIDType_POSTAL_ADDRESS,
	"PHONE_NUMBER":                 enum.PartySubIDType_PHONE_NUMBER,
	"EMAIL_ADDRESS":                enum.PartySubIDType_EMAIL_ADDRESS,
	"CONTACT_NAME":                 enum.PartySubIDType_CONTACT_NAME,
}

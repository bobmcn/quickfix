//Package quotecancel msg type = Z.
package quotecancel

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a QuoteCancel wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//QuoteReqID is a non-required field for QuoteCancel.
func (m Message) QuoteReqID() (*field.QuoteReqIDField, quickfix.MessageRejectError) {
	f := &field.QuoteReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteReqID reads a QuoteReqID from QuoteCancel.
func (m Message) GetQuoteReqID(f *field.QuoteReqIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a required field for QuoteCancel.
func (m Message) QuoteID() (*field.QuoteIDField, quickfix.MessageRejectError) {
	f := &field.QuoteIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from QuoteCancel.
func (m Message) GetQuoteID(f *field.QuoteIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteCancelType is a required field for QuoteCancel.
func (m Message) QuoteCancelType() (*field.QuoteCancelTypeField, quickfix.MessageRejectError) {
	f := &field.QuoteCancelTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteCancelType reads a QuoteCancelType from QuoteCancel.
func (m Message) GetQuoteCancelType(f *field.QuoteCancelTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteResponseLevel is a non-required field for QuoteCancel.
func (m Message) QuoteResponseLevel() (*field.QuoteResponseLevelField, quickfix.MessageRejectError) {
	f := &field.QuoteResponseLevelField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteResponseLevel reads a QuoteResponseLevel from QuoteCancel.
func (m Message) GetQuoteResponseLevel(f *field.QuoteResponseLevelField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for QuoteCancel.
func (m Message) NoPartyIDs() (*field.NoPartyIDsField, quickfix.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from QuoteCancel.
func (m Message) GetNoPartyIDs(f *field.NoPartyIDsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for QuoteCancel.
func (m Message) Account() (*field.AccountField, quickfix.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from QuoteCancel.
func (m Message) GetAccount(f *field.AccountField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for QuoteCancel.
func (m Message) AccountType() (*field.AccountTypeField, quickfix.MessageRejectError) {
	f := &field.AccountTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from QuoteCancel.
func (m Message) GetAccountType(f *field.AccountTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for QuoteCancel.
func (m Message) TradingSessionID() (*field.TradingSessionIDField, quickfix.MessageRejectError) {
	f := &field.TradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from QuoteCancel.
func (m Message) GetTradingSessionID(f *field.TradingSessionIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for QuoteCancel.
func (m Message) TradingSessionSubID() (*field.TradingSessionSubIDField, quickfix.MessageRejectError) {
	f := &field.TradingSessionSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from QuoteCancel.
func (m Message) GetTradingSessionSubID(f *field.TradingSessionSubIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoQuoteEntries is a non-required field for QuoteCancel.
func (m Message) NoQuoteEntries() (*field.NoQuoteEntriesField, quickfix.MessageRejectError) {
	f := &field.NoQuoteEntriesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoQuoteEntries reads a NoQuoteEntries from QuoteCancel.
func (m Message) GetNoQuoteEntries(f *field.NoQuoteEntriesField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized MessageBuilder with specified required fields for QuoteCancel.
func New(
	quoteid *field.QuoteIDField,
	quotecanceltype *field.QuoteCancelTypeField) Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX43))
	builder.Header.Set(field.NewMsgType("Z"))
	builder.Body.Set(quoteid)
	builder.Body.Set(quotecanceltype)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX43, "Z", r
}

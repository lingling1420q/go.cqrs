package ycq

import (
	. "gopkg.in/check.v1"
)

var _ = Suite(&InternalEventBusSuite{})

type InternalEventBusSuite struct {
	bus *InternalEventBus
}

func (s *InternalEventBusSuite) SetUpTest(c *C) {
	s.bus = NewInternalEventBus()
}

func (s *InternalEventBusSuite) Test_NewHandlerEventBus(c *C) {
	bus := NewInternalEventBus()
	c.Assert(bus, NotNil)
}

func (s *InternalEventBusSuite) TestEventBusPublishesEventsToHandlers(c *C) {
	h := NewMockEventHandler()
	ev := NewTestEventMessage(yooid())
	s.bus.AddHandler(h, &SomeEvent{})

	s.bus.PublishEvent(ev)

	c.Assert(h.events[0], Equals, ev)
}

func (s *InternalEventBusSuite) TestEventBusPublishesEventsToLocalHandlers(c *C) {
	l := NewMockEventHandler()
	ev := NewTestEventMessage(yooid())
	s.bus.AddLocalHandler(l)

	s.bus.PublishEvent(ev)

	c.Assert(l.events[0], Equals, ev)
}

func (s *InternalEventBusSuite) TestEventBusPublishesEventsToGlobalHandlers(c *C) {
	g := NewMockEventHandler()
	ev := NewTestEventMessage(yooid())
	s.bus.AddGlobalHandler(g)

	s.bus.PublishEvent(ev)

	c.Assert(g.events[0], Equals, ev)
}

func (s *InternalEventBusSuite) TestRegisterMultipleEventsForHandler(c *C) {
	h := NewMockEventHandler()
	ev1 := NewEventMessage(yooid(), &SomeEvent{Item: "Some Item", Count: 3456})
	ev2 := NewEventMessage(yooid(), &SomeOtherEvent{OrderID: yooid()})

	s.bus.AddHandler(h, &SomeEvent{}, &SomeOtherEvent{})

	s.bus.PublishEvent(ev1)
	s.bus.PublishEvent(ev2)

	c.Assert(h.events[0], Equals, ev1)
	c.Assert(h.events[1], Equals, ev2)
}

// Stubs

type MockEventBus struct {
	events []EventMessage
}

func (m *MockEventBus) PublishEvent(event EventMessage) {
	m.events = append(m.events, event)
}

func (m *MockEventBus) AddHandler(handler EventHandler, event ...interface{}) {}
func (m *MockEventBus) AddLocalHandler(handler EventHandler)                  {}
func (m *MockEventBus) AddGlobalHandler(handler EventHandler)                 {}
// Copyright (c) 2014 - The Event Horizon authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package eventhorizon

import (
	"context"
	"reflect"
	"runtime"
	"strings"
)

// EventHandlerType is the type of an event handler, used as its unique identifier.
type EventHandlerType string

// EventHandler is a handler of events. If registered on a bus as a handler only
// one handler of the same type will receive each event. If registered on a bus
// as an observer all handlers of the same type will receive each event.
type EventHandler interface {
	// HandlerType is the type of the handler.
	HandlerType() EventHandlerType

	// HandleEvent handles an event.
	HandleEvent(context.Context, Event) error
}

// EventHandlerFunc is a function that can be used as a event handler.
type EventHandlerFunc func(context.Context, Event) error

// HandleEvent implements the HandleEvent method of the EventHandler.
func (f EventHandlerFunc) HandleEvent(ctx context.Context, e Event) error {
	return f(ctx, e)
}

// HandlerType implements the HandlerType method of the EventHandler by returning
// the name of the package and function:
// "github.com/looplab/eventhorizon.Function" becomes "eventhorizon-Function"
func (f EventHandlerFunc) HandlerType() EventHandlerType {
	fullName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name() // Extract full func name: github.com/...
	parts := strings.Split(fullName, "/")                              // Split URL.
	name := parts[len(parts)-1]                                        // Take only the last part: package.Function.
	return EventHandlerType(strings.ReplaceAll(name, ".", "-"))        // Use - as separator.
}

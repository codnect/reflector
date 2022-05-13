package reflector

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestTypeOfChan(t *testing.T) {
	typ := TypeOf[chan string]()
	assert.True(t, IsChan(typ))
	assert.Equal(t, "chan string", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	chanType, isChan := ToChan(typ)

	assert.NotNil(t, chanType)
	assert.True(t, isChan)

	assert.False(t, chanType.CanSet())

	value, err := chanType.Value()
	assert.Empty(t, value)
	assert.NotNil(t, err)

	assert.Equal(t, BOTH, chanType.Direction())

	err = chanType.SetValue(make(chan string))
	assert.NotNil(t, err)

	chanElement := chanType.Elem()
	assert.NotNil(t, chanElement)
	assert.Equal(t, "string", chanElement.Name())

	chanElementString, isString := ToString(chanElement)
	assert.NotNil(t, chanElementString)
	assert.True(t, isString)

	err = chanType.Send("test-value")
	assert.NotNil(t, err)

	receivedValue, err := chanType.Receive()
	assert.Empty(t, receivedValue)
	assert.NotNil(t, err)

	capacity, err := chanType.Cap()
	assert.Equal(t, -1, capacity)
	assert.NotNil(t, err)

	newChan, _ := chanType.Instantiate()
	assert.NotNil(t, newChan)

	chanPtrVal, ok := newChan.Val().(*chan string)
	assert.True(t, ok)
	assert.Empty(t, *chanPtrVal)

	chanVal, ok := newChan.Elem().(chan string)
	assert.True(t, ok)
	assert.Empty(t, chanVal)
}

func TestTypeOfChanPointer(t *testing.T) {
	ptrType := TypeOf[*<-chan string]()
	assert.True(t, IsPointer(ptrType))
	ptr, isPtr := ToPointer(ptrType)

	assert.True(t, isPtr)
	assert.NotNil(t, ptr)

	assert.Equal(t, "*<-chan string", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.Nil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsChan(typ))
	assert.Equal(t, "<-chan string", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.False(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.Nil(t, typ.ReflectValue())

	chanType, isChan := ToChan(typ)

	assert.NotNil(t, chanType)
	assert.True(t, isChan)

	assert.False(t, chanType.CanSet())

	value, err := chanType.Value()
	assert.Empty(t, value)
	assert.NotNil(t, err)

	assert.Equal(t, RECEIVE, chanType.Direction())

	err = chanType.SetValue(make(chan string))
	assert.NotNil(t, err)

	chanElement := chanType.Elem()
	assert.NotNil(t, chanElement)
	assert.Equal(t, "string", chanElement.Name())

	chanElementString, isString := ToString(chanElement)
	assert.NotNil(t, chanElementString)
	assert.True(t, isString)

	err = chanType.Send("test-value")
	assert.NotNil(t, err)

	receivedValue, err := chanType.Receive()
	assert.Empty(t, receivedValue)
	assert.NotNil(t, err)

	capacity, err := chanType.Cap()
	assert.Equal(t, -1, capacity)
	assert.NotNil(t, err)

	newChan, _ := chanType.Instantiate()
	assert.NotNil(t, newChan)

	chanPtrVal, ok := newChan.Val().(*<-chan string)
	assert.True(t, ok)
	assert.Empty(t, *chanPtrVal)

	chanVal, ok := newChan.Elem().(<-chan string)
	assert.True(t, ok)
	assert.Empty(t, chanVal)
}

func TestTypeOfChanObject(t *testing.T) {
	val := make(chan string, 10)

	typ := TypeOfAny(val)
	assert.True(t, IsChan(typ))
	assert.Equal(t, "chan string", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	chanType, isChan := ToChan(typ)

	assert.NotNil(t, chanType)
	assert.True(t, isChan)

	assert.False(t, chanType.CanSet())

	value, err := chanType.Value()
	assert.Empty(t, value)
	assert.Nil(t, err)

	assert.Equal(t, BOTH, chanType.Direction())

	err = chanType.SetValue(make(chan string))
	assert.NotNil(t, err)

	chanElement := chanType.Elem()
	assert.NotNil(t, chanElement)
	assert.Equal(t, "string", chanElement.Name())

	chanElementString, isString := ToString(chanElement)
	assert.NotNil(t, chanElementString)
	assert.True(t, isString)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		receivedValue, err := chanType.Receive()
		assert.NotEmpty(t, receivedValue)
		assert.Nil(t, err)
		wg.Done()
	}()

	err = chanType.Send("test-value1")
	assert.Nil(t, err)

	wg.Wait()

	wg.Add(1)
	go func() {
		receivedValue, err := chanType.Receive()
		assert.NotEmpty(t, receivedValue)
		assert.Nil(t, err)
		wg.Done()
	}()

	time.Sleep(time.Second * 2)
	err = chanType.TrySend("test-value2")
	assert.Nil(t, err)

	wg.Wait()

	wg.Add(1)
	go func() {
		err = chanType.Send("test-value3")
		wg.Done()
	}()

	time.Sleep(time.Second * 2)
	receivedValue, err := chanType.TryReceive()
	assert.NotEmpty(t, receivedValue)
	assert.Nil(t, err)

	wg.Wait()

	capacity, err := chanType.Cap()
	assert.Equal(t, 10, capacity)
	assert.Nil(t, err)

	newChan, _ := chanType.Instantiate()
	assert.NotNil(t, newChan)

	chanPtrVal, ok := newChan.Val().(*chan string)
	assert.True(t, ok)
	assert.Empty(t, *chanPtrVal)

	chanVal, ok := newChan.Elem().(chan string)
	assert.True(t, ok)
	assert.Empty(t, chanVal)
}

func TestTypeOfChanObjectPointer(t *testing.T) {
	val := make(chan string)

	ptrType := TypeOfAny(&val)
	assert.True(t, IsPointer(ptrType))
	ptr, isPtr := ToPointer(ptrType)

	assert.True(t, isPtr)
	assert.NotNil(t, ptr)

	assert.Equal(t, "*chan string", ptr.Name())
	assert.Equal(t, "", ptr.PackageName())
	assert.NotNil(t, ptr.ReflectType())
	assert.NotNil(t, ptr.ReflectValue())

	typ := ptr.Elem()

	assert.True(t, IsChan(typ))
	assert.Equal(t, "chan string", typ.Name())
	assert.Equal(t, "", typ.PackageName())

	assert.True(t, typ.HasValue())
	assert.NotNil(t, typ.ReflectType())
	assert.NotNil(t, typ.ReflectValue())

	chanType, isChan := ToChan(typ)

	assert.NotNil(t, chanType)
	assert.True(t, isChan)

	assert.True(t, chanType.CanSet())

	value, err := chanType.Value()
	assert.Empty(t, value)
	assert.Nil(t, err)

	assert.Equal(t, BOTH, chanType.Direction())

	err = chanType.SetValue(make(chan string))
	assert.Nil(t, err)

	chanElement := chanType.Elem()
	assert.NotNil(t, chanElement)
	assert.Equal(t, "string", chanElement.Name())

	chanElementString, isString := ToString(chanElement)
	assert.NotNil(t, chanElementString)
	assert.True(t, isString)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		receivedValue, err := chanType.Receive()
		assert.NotEmpty(t, receivedValue)
		assert.Nil(t, err)
		wg.Done()
	}()

	err = chanType.Send("test-value1")
	assert.Nil(t, err)

	wg.Wait()

	wg.Add(1)
	go func() {
		receivedValue, err := chanType.Receive()
		assert.NotEmpty(t, receivedValue)
		assert.Nil(t, err)
		wg.Done()
	}()

	time.Sleep(time.Second * 2)
	err = chanType.TrySend("test-value2")
	assert.Nil(t, err)

	wg.Wait()

	wg.Add(1)
	go func() {
		err = chanType.Send("test-value3")
		wg.Done()
	}()

	time.Sleep(time.Second * 2)
	receivedValue, err := chanType.TryReceive()
	assert.NotEmpty(t, receivedValue)
	assert.Nil(t, err)

	wg.Wait()

	capacity, err := chanType.Cap()
	assert.Equal(t, 0, capacity)
	assert.Nil(t, err)

	newChan, _ := chanType.Instantiate()
	assert.NotNil(t, newChan)

	chanPtrVal, ok := newChan.Val().(*chan string)
	assert.True(t, ok)
	assert.Empty(t, *chanPtrVal)

	chanVal, ok := newChan.Elem().(chan string)
	assert.True(t, ok)
	assert.Empty(t, chanVal)
}

/**
 * MIT License
 *
 * Copyright (c) 2018 CNES
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
package consumer

import (
	. "github.com/ccsdsmo/malgo/com"
	. "github.com/ccsdsmo/malgo/mal"
	. "github.com/ccsdsmo/malgo/mal/api"

	. "github.com/EtienneLndr/archiveservice/archive/constants"
	. "github.com/EtienneLndr/archiveservice/data"
)

// InvokeConsumer :
type InvokeConsumer struct {
	ctx     *Context
	cctx    *ClientContext
	op      InvokeOperation
	factory EncodingFactory
}

// ProgressConsumer :
type ProgressConsumer struct {
	ctx     *Context
	cctx    *ClientContext
	op      ProgressOperation
	factory EncodingFactory
}

// RequestConsumer :
type RequestConsumer struct {
	ctx     *Context
	cctx    *ClientContext
	op      RequestOperation
	factory EncodingFactory
}

// SubmitConsumer :
type SubmitConsumer struct {
	ctx     *Context
	cctx    *ClientContext
	op      SubmitOperation
	factory EncodingFactory
}

// Close :
func (i *InvokeConsumer) Close() {
	i.ctx.Close()
}

// Close :
func (p *ProgressConsumer) Close() {
	p.ctx.Close()
}

// Close :
func (r *RequestConsumer) Close() {
	r.ctx.Close()
}

// Close :
func (s *SubmitConsumer) Close() {
	s.ctx.Close()
}

//======================================================================//
//								CONSUMERS								//
//======================================================================//
// Create a consumer for an invoke operation
func createInvokeConsumer(url string, factory EncodingFactory, providerURI *URI, typeOfConsumer string, operation UShort) (*InvokeConsumer, error) {
	ctx, err := NewContext(url)
	if err != nil {
		return nil, err
	}

	cctx, err := NewClientContext(ctx, typeOfConsumer)
	if err != nil {
		return nil, err
	}

	op := cctx.NewInvokeOperation(providerURI,
		SERVICE_AREA_NUMBER,
		SERVICE_AREA_VERSION,
		ARCHIVE_SERVICE_SERVICE_NUMBER,
		operation)

	consumer := &InvokeConsumer{ctx, cctx, op, factory}

	return consumer, nil
}

// Create a consumer for a progress operation
func createProgressConsumer(url string, factory EncodingFactory, providerURI *URI, typeOfConsumer string, operation UShort) (*ProgressConsumer, error) {
	ctx, err := NewContext(url)
	if err != nil {
		return nil, err
	}

	cctx, err := NewClientContext(ctx, typeOfConsumer)
	if err != nil {
		return nil, err
	}

	op := cctx.NewProgressOperation(providerURI,
		SERVICE_AREA_NUMBER,
		SERVICE_AREA_VERSION,
		ARCHIVE_SERVICE_SERVICE_NUMBER,
		operation)

	consumer := &ProgressConsumer{ctx, cctx, op, factory}

	return consumer, nil
}

// Create a consumer for a request operation
func createRequestConsumer(url string, factory EncodingFactory, providerURI *URI, typeOfConsumer string, operation UShort) (*RequestConsumer, error) {
	ctx, err := NewContext(url)
	if err != nil {
		return nil, err
	}

	cctx, err := NewClientContext(ctx, typeOfConsumer)
	if err != nil {
		return nil, err
	}

	op := cctx.NewRequestOperation(providerURI,
		SERVICE_AREA_NUMBER,
		SERVICE_AREA_VERSION,
		ARCHIVE_SERVICE_SERVICE_NUMBER,
		operation)

	consumer := &RequestConsumer{ctx, cctx, op, factory}

	return consumer, nil
}

// Create a consumer for a submit operation
func createSubmitConsumer(url string, factory EncodingFactory, providerURI *URI, typeOfConsumer string, operation UShort) (*SubmitConsumer, error) {
	ctx, err := NewContext(url)
	if err != nil {
		return nil, err
	}

	cctx, err := NewClientContext(ctx, typeOfConsumer)
	if err != nil {
		return nil, err
	}

	op := cctx.NewSubmitOperation(providerURI,
		SERVICE_AREA_NUMBER,
		SERVICE_AREA_VERSION,
		ARCHIVE_SERVICE_SERVICE_NUMBER,
		operation)

	consumer := &SubmitConsumer{ctx, cctx, op, factory}

	return consumer, nil
}

//======================================================================//
//								RETRIEVE								//
//======================================================================//
// StartRetrieveConsumer : TODO
func StartRetrieveConsumer(url string, factory EncodingFactory, providerURI *URI, objectType ObjectType, identifierList IdentifierList, longList LongList) (*InvokeConsumer, *ArchiveDetailsList, ElementList, error) {
	// Create the consumer
	consumer, err := createInvokeConsumer(url, factory, providerURI, "consumerRetrieve", OPERATION_IDENTIFIER_RETRIEVE)
	if err != nil {
		return nil, nil, nil, err
	}

	// Call Invoke operation
	err = consumer.retrieveInvoke(objectType, identifierList, longList)
	if err != nil {
		return nil, nil, nil, err
	}

	// Call Response operation
	archiveDetailsList, elementList, err := consumer.retrieveResponse()
	if err != nil {
		return nil, nil, nil, err
	}

	return consumer, archiveDetailsList, elementList, nil
}

// Invoke & Ack
func (consumer *InvokeConsumer) retrieveInvoke(objectType ObjectType, identifierList IdentifierList, longList LongList) error {
	// Create the encoder
	encoder := consumer.factory.NewEncoder(make([]byte, 0, 8192))
	// Encode ObjectType
	err := objectType.Encode(encoder)
	if err != nil {
		return err
	}

	// Encode IdentifierList
	err = identifierList.Encode(encoder)
	if err != nil {
		return err
	}

	// Encode LongList
	err = longList.Encode(encoder)
	if err != nil {
		return err
	}

	// Call Invoke operation
	_, err = consumer.op.Invoke(encoder.Body())
	if err != nil {
		return err
	}

	return nil
}

// Response
func (consumer *InvokeConsumer) retrieveResponse() (*ArchiveDetailsList, ElementList, error) {
	// Call Response operation
	resp, err := consumer.op.GetResponse()
	if err != nil {
		return nil, nil, err
	}

	// Create the decoder
	decoder := consumer.factory.NewDecoder(resp.Body)

	// Decode ArchiveDetailsList
	archiveDetails, err := decoder.DecodeElement(NullArchiveDetailsList)
	if err != nil {
		return nil, nil, err
	}

	// Decode ElementList
	elementList, err := decoder.DecodeAbstractElement()
	if err != nil {
		return nil, nil, err
	}

	return archiveDetails.(*ArchiveDetailsList), elementList.(ElementList), nil
}

//======================================================================//
//								QUERY									//
//======================================================================//
// StartQueryConsumer : TODO
func StartQueryConsumer(url string, factory EncodingFactory, providerURI *URI, boolean Boolean, objectType ObjectType, archiveQueryList ArchiveQueryList, queryFilterList QueryFilterList) (*ProgressConsumer, []interface{}, error) {
	// Create the consumer
	consumer, err := createProgressConsumer(url, factory, providerURI, "consumerQuery", OPERATION_IDENTIFIER_QUERY)
	if err != nil {
		return nil, nil, err
	}

	// Call Progress function
	err = consumer.queryProgress(boolean, objectType, archiveQueryList, queryFilterList)
	if err != nil {
		return nil, nil, err
	}

	// Create the interface that will receive all the responses
	responses := []interface{}{}
	// Call Update operation
	respObjType, respIDList, respArchDetList, respElemList, err := consumer.queryUpdate()
	if err != nil {
		return nil, nil, err
	}

	for (respObjType != nil) && (respIDList != nil) && (respArchDetList != nil) && (respElemList != nil) {
		// Put the objects in the interface
		responses = append(responses, respObjType, respIDList, respArchDetList, respElemList)

		// Call Update operation until it returns nil variables
		respObjType, respIDList, respArchDetList, respElemList, err = consumer.queryUpdate()
		if err != nil {
			return nil, nil, err
		}
	}

	// Call Response operation
	respObjType, respIDList, respArchDetList, respElemList, err = consumer.queryResponse()
	if err != nil {
		return nil, nil, err
	}

	// Put the objects in the interface
	responses = append(responses, respObjType, respIDList, respArchDetList, respElemList)

	return consumer, responses, nil
}

// Progress & Ack
func (consumer *ProgressConsumer) queryProgress(boolean Boolean, objectType ObjectType, archiveQueryList ArchiveQueryList, queryFilterList QueryFilterList) error {
	// Create the encoder
	encoder := consumer.factory.NewEncoder(make([]byte, 0, 8192))

	// Encode Boolean
	err := boolean.Encode(encoder)
	if err != nil {
		return err
	}

	// Encode ObjectType
	err = objectType.Encode(encoder)
	if err != nil {
		return err
	}

	// Encode ArchiveQueryList
	err = archiveQueryList.Encode(encoder)
	if err != nil {
		return err
	}

	// Encode QueryFilterList
	err = encoder.EncodeAbstractElement(queryFilterList)
	if err != nil {
		return err
	}

	// Call Progress operation
	_, err = consumer.op.Progress(encoder.Body())
	if err != nil {
		return err
	}

	return nil
}

// Update
func (consumer *ProgressConsumer) queryUpdate() (*ObjectType, *IdentifierList, *ArchiveDetailsList, ElementList, error) {
	// Call Update operation
	updt, err := consumer.op.GetUpdate()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	if updt != nil {
		// Create the decoder to decode the multiple variables
		decoder := consumer.factory.NewDecoder(updt.Body)

		// Decode ObjectType
		objectType, err := decoder.DecodeElement(NullObjectType)
		if err != nil {
			return nil, nil, nil, nil, err
		}

		// Decode IdentifierList
		identifierList, err := decoder.DecodeElement(NullIdentifierList)
		if err != nil {
			return nil, nil, nil, nil, err
		}

		// Decode ArchiveDetailsList
		archiveDetailsList, err := decoder.DecodeElement(NullArchiveQueryList)
		if err != nil {
			return nil, nil, nil, nil, err
		}

		// Decode ElementList
		elementList, err := decoder.DecodeAbstractElement()
		if err != nil {
			return nil, nil, nil, nil, err
		}

		return objectType.(*ObjectType), identifierList.(*IdentifierList), archiveDetailsList.(*ArchiveDetailsList), elementList.(ElementList), nil
	}
	return nil, nil, nil, nil, nil
}

// Response
func (consumer *ProgressConsumer) queryResponse() (*ObjectType, *IdentifierList, *ArchiveDetailsList, ElementList, error) {
	// Call Update operation
	resp, err := consumer.op.GetResponse()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	// Create the decoder to decode the multiple variables
	decoder := consumer.factory.NewDecoder(resp.Body)

	// Decode ObjectType
	objectType, err := decoder.DecodeElement(NullObjectType)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	// Decode IdentifierList
	identifierList, err := decoder.DecodeElement(NullIdentifierList)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	// Decode ArchiveDetailsList
	archiveDetailsList, err := decoder.DecodeElement(NullArchiveQueryList)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	// Decode ElementList
	elementList, err := decoder.DecodeAbstractElement()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return objectType.(*ObjectType), identifierList.(*IdentifierList), archiveDetailsList.(*ArchiveDetailsList), elementList.(ElementList), nil
}

//======================================================================//
//								COUNT									//
//======================================================================//
// StartCountConsumer : TODO
func StartCountConsumer(url string, factory EncodingFactory, providerURI *URI, objectType ObjectType, archiveQueryList ArchiveQueryList, queryFilterList QueryFilterList) (*InvokeConsumer, *LongList, error) {
	// Create the consumer
	consumer, err := createInvokeConsumer(url, factory, providerURI, "consumerCount", OPERATION_IDENTIFIER_COUNT)
	if err != nil {
		return nil, nil, err
	}

	// Call Invoke function
	err = consumer.countInvoke(objectType, archiveQueryList, queryFilterList)
	if err != nil {
		return nil, nil, err
	}

	// Call Response function
	longList, err := consumer.countResponse()
	if err != nil {
		return nil, nil, err
	}

	return consumer, longList, nil
}

// Invoke & Ack
func (consumer *InvokeConsumer) countInvoke(objectType ObjectType, archiveQueryList ArchiveQueryList, queryFilterList QueryFilterList) error {
	// Create the encoder
	encoder := consumer.factory.NewEncoder(make([]byte, 0, 8192))

	// Encode ObjectType
	err := objectType.Encode(encoder)
	if err != nil {
		return err
	}

	// Encode ArchiveQueryList
	err = archiveQueryList.Encode(encoder)
	if err != nil {
		return err
	}

	// Encode QueryFilterList
	err = encoder.EncodeAbstractElement(queryFilterList)
	if err != nil {
		return err
	}

	// Call Invoke operation
	// TODO: we should retrieve the msg to verify if the ack is an error or not
	_, err = consumer.op.Invoke(encoder.Body())
	if err != nil {
		return err
	}

	return nil
}

// Response
func (consumer *InvokeConsumer) countResponse() (*LongList, error) {
	// Call Response operation
	resp, err := consumer.op.GetResponse()
	if err != nil {
		return nil, err
	}

	// Create the decoder
	decoder := consumer.factory.NewDecoder(resp.Body)

	// Decode LongList
	longList, err := decoder.DecodeElement(NullLongList)
	if err != nil {
		return nil, err
	}

	return longList.(*LongList), nil
}

//======================================================================//
//								STORE									//
//======================================================================//
// StartStoreConsumer : TODO
func StartStoreConsumer(url string, factory EncodingFactory, providerURI *URI, boolean Boolean, objectType ObjectType, identifierList IdentifierList, archiveDetailsList ArchiveDetailsList, elementList ElementList) (*RequestConsumer, *LongList, error) {
	// Create the consumer
	consumer, err := createRequestConsumer(url, factory, providerURI, "consumerStore", OPERATION_IDENTIFIER_STORE)
	if err != nil {
		return nil, nil, err
	}

	// Call Request function and retrieve the Response
	longList, err := consumer.storeRequest(boolean, objectType, identifierList, archiveDetailsList, elementList)
	if err != nil {
		return nil, nil, err
	}

	return consumer, longList, nil
}

// Request & Response
func (consumer *RequestConsumer) storeRequest(boolean Boolean, objectType ObjectType, identifierList IdentifierList, archiveDetailsList ArchiveDetailsList, elementList ElementList) (*LongList, error) {
	// Create the encoder
	encoder := consumer.factory.NewEncoder(make([]byte, 0, 8192))

	// Encode Boolean
	err := boolean.Encode(encoder)
	if err != nil {
		return nil, err
	}

	// Encode ObjectType
	err = objectType.Encode(encoder)
	if err != nil {
		return nil, err
	}

	// Encode IdentifierList
	err = identifierList.Encode(encoder)
	if err != nil {
		return nil, err
	}

	// Encode ArchiveDetailsList
	err = archiveDetailsList.Encode(encoder)
	if err != nil {
		return nil, err
	}

	// Encode ElementList
	err = encoder.EncodeAbstractElement(elementList)
	if err != nil {
		return nil, err
	}

	// Call Request operation and retrieve the Response
	resp, err := consumer.op.Request(encoder.Body())
	if err != nil {
		return nil, err
	}

	// Create the decoder
	decoder := consumer.factory.NewDecoder(resp.Body)

	// Decode LongList
	longList, err := decoder.DecodeElement(NullLongList)
	if err != nil {
		return nil, err
	}

	return longList.(*LongList), nil
}

//======================================================================//
//								UPDATE									//
//======================================================================//
// StartUpdateConsumer : TODO
func StartUpdateConsumer(url string, factory EncodingFactory, providerURI *URI, objectType ObjectType, identifierList IdentifierList, archiveDetailsList ArchiveDetailsList, elementList ElementList) (*SubmitConsumer, error) {
	// Create the consumer
	consumer, err := createSubmitConsumer(url, factory, providerURI, "consumerUpdate", OPERATION_IDENTIFIER_UPDATE)
	if err != nil {
		return nil, err
	}

	// Call Submit function
	err = consumer.updateSubmit(objectType, identifierList, archiveDetailsList, elementList)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}

// Submit & Ack
func (consumer *SubmitConsumer) updateSubmit(objectType ObjectType, identifierList IdentifierList, archiveDetailsList ArchiveDetailsList, elementList ElementList) error {
	// Create the encoder
	encoder := consumer.factory.NewEncoder(make([]byte, 0, 8192))

	// Encode ObjectType
	err := objectType.Encode(encoder)
	if err != nil {
		return err
	}

	// Encode IdentifierList
	err = identifierList.Encode(encoder)
	if err != nil {
		return err
	}

	// Encode ArchiveDetailsList
	err = archiveDetailsList.Encode(encoder)
	if err != nil {
		return err
	}

	// Encode ElementList
	err = encoder.EncodeAbstractElement(elementList)
	if err != nil {
		return err
	}

	// Call Submit operation
	_, err = consumer.op.Submit(encoder.Body())
	if err != nil {
		return err
	}

	return nil
}

//======================================================================//
//								DELETE									//
//======================================================================//
// StartDeleteConsumer : TODO
func StartDeleteConsumer(url string, factory EncodingFactory, providerURI *URI, objectType ObjectType, identifierList IdentifierList, longList LongList) (*RequestConsumer, *LongList, error) {
	// Create the consumer
	consumer, err := createRequestConsumer(url, factory, providerURI, "consumerDelete", OPERATION_IDENTIFIER_DELETE)
	if err != nil {
		return nil, nil, err
	}

	// Call Request function and retrieve the Response
	respLongList, err := consumer.deleteRequest(objectType, identifierList, longList)
	if err != nil {
		return nil, nil, err
	}

	return consumer, respLongList, nil
}

// Request & Reponse
func (consumer *RequestConsumer) deleteRequest(objectType ObjectType, identifierList IdentifierList, longList LongList) (*LongList, error) {
	// Create the encoder
	encoder := consumer.factory.NewEncoder(make([]byte, 0, 8192))

	// Encode ObjectType
	err := objectType.Encode(encoder)
	if err != nil {
		return nil, err
	}

	// Encode IdentifierList
	err = identifierList.Encode(encoder)
	if err != nil {
		return nil, err
	}

	// Encode LongList
	err = longList.Encode(encoder)
	if err != nil {
		return nil, err
	}

	// Call Request operation and retrieve the Response
	resp, err := consumer.op.Request(encoder.Body())
	if err != nil {
		return nil, err
	}

	// Create the decoder
	decoder := consumer.factory.NewDecoder(resp.Body)

	// Decode LongList
	respLongList, err := decoder.DecodeElement(NullLongList)
	if err != nil {
		return nil, err
	}

	return respLongList.(*LongList), nil
}

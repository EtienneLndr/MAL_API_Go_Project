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
package archive

import (
	"errors"
	"fmt"
	"time"

	. "github.com/ccsdsmo/malgo/com"
	. "github.com/ccsdsmo/malgo/mal"
	. "github.com/ccsdsmo/malgo/mal/encoding/binary"
	. "github.com/etiennelndr/archiveservice/archive/constants"
	. "github.com/etiennelndr/archiveservice/archive/consumer"
	. "github.com/etiennelndr/archiveservice/archive/provider"
	. "github.com/etiennelndr/archiveservice/data"
	. "github.com/etiennelndr/archiveservice/service"
)

type ArchiveService struct {
	areaIdentifier    Identifier
	serviceIdentifier Identifier
	areaNumber        Integer
	serviceNumber     Integer
	areaVersion       Integer
}

// Constants for the providers and consumers
const (
	providerURL         = "maltcp://127.0.0.1:12400"
	providerURLRetrieve = providerURL + "/providerRetrieve"
	providerURLQuery    = providerURL + "/providerQuery"
	providerURLCount    = providerURL + "/providerCount"
	providerURLStore    = providerURL + "/providerStore"
	providerURLUpdate   = providerURL + "/providerUpdate"
	providerURLDelete   = providerURL + "/providerDelete"
	consumerURL         = "maltcp://127.0.0.1:14200"
	/*consumerURLRetrieve = consumerURL + "/consumerRetrieve"
	consumerURLQuery    = consumerURL + "/consumerQuery"
	consumerURLCount    = consumerURL + "/consumerCount"
	consumerURLStore    = consumerURL + "/consumerStore"
	consumerURLUpdate   = consumerURL + "/consumerUpdate"
	consumerURLDelete   = consumerURL + "/consumerDelete"*/
)

// CreateService : TODO
func (*ArchiveService) CreateService() Service {
	archiveService := &ArchiveService{
		ARCHIVE_SERVICE_AREA_IDENTIFIER,
		ARCHIVE_SERVICE_SERVICE_IDENTIFIER,
		SERVICE_AREA_NUMBER,
		ARCHIVE_SERVICE_SERVICE_NUMBER,
		SERVICE_AREA_VERSION,
	}

	return archiveService
}

//======================================================================//
//								RETRIEVE								//
//======================================================================//
/**
 * Operation        : Retrieve
 * Operation number : 1
 */
func (archiveService *ArchiveService) retrieveProvider() (*Provider, error) {
	// Maybe we should not have to return an error
	fmt.Println("Creation : Retrieve Provider")

	transport := new(FixedBinaryEncoding)
	provider, err := StartRetrieveProvider(providerURL, transport)
	if err != nil {
		return nil, err
	}

	return provider, nil
}

func (archiveService *ArchiveService) retrieveConsumer(objectType ObjectType, identifierList IdentifierList, longList LongList) (*InvokeConsumer, *ArchiveDetailsList, ElementList, error) {
	// Maybe we should not have to return an error
	fmt.Println("Creation : Retrieve Consumer")

	// IN
	transport := new(FixedBinaryEncoding)
	var providerURI = NewURI(providerURLRetrieve)
	// OUT
	consumer, archiveDetailsList, elementList, err := StartRetrieveConsumer(consumerURL,
		transport,
		providerURI,
		objectType,
		identifierList,
		longList)
	if err != nil {
		return nil, nil, nil, err
	}

	return consumer, archiveDetailsList, elementList, nil
}

//======================================================================//
//								QUERY									//
//======================================================================//
/**
 * Operation        : Query
 * Operation number : 2
 */
func (archiveService *ArchiveService) queryProvider() (*Provider, error) {
	// Maybe we should not have to return an error
	fmt.Println("Creation : Query Provider")

	transport := new(FixedBinaryEncoding)
	provider, err := StartQueryProvider(providerURL, transport)
	if err != nil {
		return nil, err
	}

	return provider, nil
}

func (archiveService *ArchiveService) queryConsumer(boolean Boolean, objectType ObjectType, archiveQueryList ArchiveQueryList, queryFilterList QueryFilterList) (*ProgressConsumer, []interface{}, error) {
	// Maybe we should not have to return an error
	fmt.Println("Creation : Retrieve Consumer")

	// IN
	transport := new(FixedBinaryEncoding)
	var providerURI = NewURI(providerURLQuery)
	// OUT
	consumer, responses, err := StartQueryConsumer(consumerURL,
		transport,
		providerURI,
		boolean,
		objectType,
		archiveQueryList,
		queryFilterList)
	if err != nil {
		return nil, nil, err
	}

	return consumer, responses, nil
}

//======================================================================//
//								COUNT									//
//======================================================================//
/**
 * Operation        : Count
 * Operation number : 3
 */
func (archiveService *ArchiveService) countProvider() (*Provider, error) {
	// Maybe we should not have to return an error
	fmt.Println("Creation : Count Provider")

	transport := new(FixedBinaryEncoding)
	provider, err := StartCountProvider(providerURL, transport)
	if err != nil {
		return nil, err
	}

	return provider, nil
}

func (archiveService *ArchiveService) countConsumer(objectType ObjectType, archiveQueryList ArchiveQueryList, queryFilterList QueryFilterList) (*InvokeConsumer, *LongList, error) {
	// Maybe we should not have to return an error
	fmt.Println("Creation : Retrieve Consumer")

	// IN
	transport := new(FixedBinaryEncoding)
	var providerURI = NewURI(providerURLCount)
	// OUT
	consumer, longList, err := StartCountConsumer(consumerURL,
		transport,
		providerURI,
		objectType,
		archiveQueryList,
		queryFilterList)
	if err != nil {
		return nil, nil, err
	}

	return consumer, longList, nil
}

//======================================================================//
//								STORE									//
//======================================================================//
/**
 * Operation        : Store
 * Operation number : 4
 */
func (archiveService *ArchiveService) storeProvider() (*Provider, error) {
	// Maybe we should not have to return an error
	fmt.Println("Creation : Store Provider")

	transport := new(FixedBinaryEncoding)
	provider, err := StartStoreProvider(providerURL, transport)
	if err != nil {
		return nil, err
	}

	return provider, nil
}

func (archiveService *ArchiveService) storeConsumer(boolean Boolean, objectType ObjectType, identifierList IdentifierList, archiveDetailsList ArchiveDetailsList, elementList ElementList) (*RequestConsumer, *LongList, error) {
	// Maybe we should not have to return an error
	fmt.Println("Creation : Retrieve Consumer")

	// IN
	transport := new(FixedBinaryEncoding)
	var providerURI = NewURI(providerURLStore)
	// OUT
	consumer, longList, err := StartStoreConsumer(consumerURL,
		transport,
		providerURI,
		boolean,
		objectType,
		identifierList,
		archiveDetailsList,
		elementList)
	if err != nil {
		return nil, nil, err
	}

	return consumer, longList, nil
}

//======================================================================//
//								UPDATE									//
//======================================================================//
/**
 * Operation        : Update
 * Operation number : 5
 */
func (archiveService *ArchiveService) updateProvider() (*Provider, error) {
	// Maybe we should not have to return an error
	fmt.Println("Creation : Update Provider")

	transport := new(FixedBinaryEncoding)
	provider, err := StartUpdateProvider(providerURL, transport)
	if err != nil {
		return nil, err
	}

	return provider, nil
}

func (archiveService *ArchiveService) updateConsumer(objectType ObjectType, identifierList IdentifierList, archiveDetailsList ArchiveDetailsList, elementList ElementList) (*SubmitConsumer, error) {
	// Maybe we should not have to return an error
	fmt.Println("Creation : Retrieve Consumer")

	// IN
	transport := new(FixedBinaryEncoding)
	var providerURI = NewURI(providerURLUpdate)
	// OUT
	consumer, err := StartUpdateConsumer(consumerURL,
		transport,
		providerURI,
		objectType,
		identifierList,
		archiveDetailsList,
		elementList)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}

//======================================================================//
//								DELETE									//
//======================================================================//
/**
 * Operation        : Delete
 * Operation number : 6
 */
func (archiveService *ArchiveService) deleteProvider() (*Provider, error) {
	// Maybe we should not have to return an error
	fmt.Println("Creation : Delete Provider")

	transport := new(FixedBinaryEncoding)
	provider, err := StartDeleteProvider(providerURL, transport)
	if err != nil {
		return nil, err
	}

	return provider, nil
}

func (archiveService *ArchiveService) deleteConsumer(objectType ObjectType, identifierList IdentifierList, longList LongList) (*RequestConsumer, *LongList, error) {
	// Maybe we should not have to return an error
	fmt.Println("Creation : Retrieve Consumer")

	// IN
	transport := new(FixedBinaryEncoding)
	var providerURI = NewURI(providerURLDelete)
	// OUT
	consumer, respLongList, err := StartDeleteConsumer(consumerURL,
		transport,
		providerURI,
		objectType,
		identifierList,
		longList)
	if err != nil {
		return nil, nil, err
	}

	return consumer, respLongList, nil
}

//======================================================================//
//							START: Consumer								//
//======================================================================//
// LaunchRetrieveConsumer : TODO
func (archiveService *ArchiveService) LaunchRetrieveConsumer(objectType ObjectType, identifierList IdentifierList, longList LongList) (*ArchiveDetailsList, ElementList, error) {
	// Start Operation
	consumer, archiveDetailsList, elementList, err := archiveService.retrieveConsumer(objectType, identifierList, longList)
	if err != nil {
		return nil, nil, err
	}

	// Close the consumer
	consumer.Close()

	return archiveDetailsList, elementList, nil
}

// LaunchQueryConsumer : TODO
func (archiveService *ArchiveService) LaunchQueryConsumer(boolean Boolean, objectType ObjectType, archiveQueryList ArchiveQueryList, queryFilterList QueryFilterList) ([]interface{}, error) {
	// Start Operation
	consumer, responses, err := archiveService.queryConsumer(boolean, objectType, archiveQueryList, queryFilterList)
	if err != nil {
		return nil, err
	}

	// Close the consumer
	consumer.Close()

	return responses, nil
}

// LaunchCountConsumer : TODO
func (archiveService *ArchiveService) LaunchCountConsumer(objectType ObjectType, archiveQueryList ArchiveQueryList, queryFilterList QueryFilterList) (*LongList, error) {
	// Start Operation
	consumer, longList, err := archiveService.countConsumer(objectType, archiveQueryList, queryFilterList)
	if err != nil {
		return nil, err
	}

	// Close the consumer
	consumer.Close()

	return longList, nil
}

// LaunchStoreConsumer : TODO
func (archiveService *ArchiveService) LaunchStoreConsumer(boolean Boolean, objectType ObjectType, identifierList IdentifierList, archiveDetailsList ArchiveDetailsList, elementList ElementList) (*LongList, error) {
	// Start Operation
	consumer, longList, err := archiveService.storeConsumer(boolean, objectType, identifierList, archiveDetailsList, elementList)
	if err != nil {
		return nil, err
	}

	// Close the consumer
	consumer.Close()

	return longList, nil
}

// LaunchUpdateConsumer : TODO
func (archiveService *ArchiveService) LaunchUpdateConsumer(objectType ObjectType, identifierList IdentifierList, archiveDetailsList ArchiveDetailsList, elementList ElementList) error {
	// Start Operation
	consumer, err := archiveService.updateConsumer(objectType, identifierList, archiveDetailsList, elementList)
	if err != nil {
		return err
	}

	// Close the consumer
	consumer.Close()

	return nil
}

// LaunchDeleteConsumer : TODO
func (archiveService *ArchiveService) LaunchDeleteConsumer(objectType ObjectType, identifierList IdentifierList, longList LongList) (*LongList, error) {
	// Start Operation
	consumer, respLongList, err := archiveService.deleteConsumer(objectType, identifierList, longList)
	if err != nil {
		return nil, err
	}

	// Close the consumer
	consumer.Close()

	return respLongList, nil
}

//======================================================================//
//							START: Provider								//
//======================================================================//
// LaunchProvider : TODO
func (archiveService *ArchiveService) LaunchProvider(operation UShort) error {
	// Declare variables
	var provider *Provider
	var err error

	// Start Operation
	switch operation {
	case OPERATION_IDENTIFIER_RETRIEVE:
		provider, err = archiveService.retrieveProvider()
		break
	case OPERATION_IDENTIFIER_QUERY:
		provider, err = archiveService.queryProvider()
		break
	case OPERATION_IDENTIFIER_COUNT:
		provider, err = archiveService.countProvider()
		break
	case OPERATION_IDENTIFIER_STORE:
		provider, err = archiveService.storeProvider()
		break
	case OPERATION_IDENTIFIER_UPDATE:
		provider, err = archiveService.updateProvider()
		break
	case OPERATION_IDENTIFIER_DELETE:
		provider, err = archiveService.deleteProvider()
		break
	default:
		return errors.New("Unknown operation")
	}

	if err != nil {
		return err
	}

	// Close the provider at the end of the function
	defer provider.Close()

	// Start communication
	var running bool = true
	for running == true {
		time.Sleep(120 * time.Second)
		running = false
	}

	return nil
}

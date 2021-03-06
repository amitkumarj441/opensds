package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/opensds/opensds/api/models"
)

// OperateVolumeReader is a Reader for the OperateVolume structure.
type OperateVolumeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OperateVolumeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewOperateVolumeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOperateVolumeOK creates a OperateVolumeOK with default headers values
func NewOperateVolumeOK() *OperateVolumeOK {
	return &OperateVolumeOK{}
}

/*OperateVolumeOK handles this case with default header values.

OK
*/
type OperateVolumeOK struct {
	Payload *models.DefaultResponse
}

func (o *OperateVolumeOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/volumes/action/{resourceType}/{id}][%d] operateVolumeOK  %+v", 200, o.Payload)
}

func (o *OperateVolumeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

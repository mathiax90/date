package date
import (
	"time"
	"strings"
	"encoding/json"
	"fmt"
    // "bytes"
    "encoding/xml"
    "errors"
    "database/sql/driver"
)

type Date struct {
	Time time.Time
}

func (d Date) MarshalJSON() ([]byte, error) {        
    return json.Marshal(d.Time.Format("2006-01-02"))
}

func (d *Date)UnmarshalJSON(in []byte) error {   
    inStr := strings.Trim(string(in), `"`)    
    var err error 
    d.Time, err = time.Parse("2006-01-02", inStr)    
    return err
}

// MarshalXML generate XML output for PrecsontructedInfo
func (d Date) MarshalXML(e *xml.Encoder, start xml.StartElement) (err error) {
    return e.EncodeElement(d.Time.Format("2006-01-02"), start)
}

func (d *Date) Scan(value interface{}) error {    
    if value == nil {
        return errors.New("Date can't Scan null use NullDate instead")
    } else {
        if DateAsString, err := driver.String.ConvertValue(value); err == nil {            
            if v, ok := DateAsString.(string); ok {
				//postresql time pattern
                d.Time, err = time.Parse("2006-01-02 15:04:05 -0700 MST", string(v))
                if err != nil {
                    return err
                }
                return nil
            }
        }
        return errors.New("failed to scan Date")
    }
}

func (d Date) String() string {
	return fmt.Sprintf("{%v}", d.Time.Format("2006-01-02"))
}
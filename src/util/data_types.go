package util

import (
	"io"
)

// ReadVarint reads one varint from the specified reader.
func ReadVarint(reader io.Reader) (int32, error) {
	bytes := 0
	result := int32(0)

	read := make([]byte, 1)

	for {
		_, err := reader.Read(read)
		if err != nil {
			return 0, err
		}

		value := read[0] & 0b01111111
		result |= int32((value << (7 * bytes)))

		bytes++

		if (read[0] & 0b10000000) == 0 {
			break
		}
	}

	return result, nil
}

// WriteVarint writes one varint to the specified writer.
func WriteVarint(value int32, writer io.Writer) error {
	for {
		temp := []byte{byte(value & 0b01111111)}
		value >>= 7

		if value != 0 {
			temp[0] |= 0b10000000
			_, err := writer.Write(temp)
			if err != nil {
				return err
			}
		} else {
			_, err := writer.Write(temp)
			if err != nil {
				return err
			}
			break
		}
	}

	return nil
}

// ReadString reads one string from the specified reader.
func ReadString(reader io.Reader) (string, error) {
	length, err := ReadVarint(reader)
	if err != nil {
		return "", err
	}

	value := make([]byte, length)
	_, err = reader.Read(value)
	if err != nil {
		return "", err
	}

	return string(value), nil
}

// WriteString writes one string to the specified writer.
func WriteString(value string, writer io.Writer) error {
	err := WriteVarint(int32(len(value)), writer)
	if err != nil {
		return err
	}

	_, err = writer.Write([]byte(value))
	if err != nil {
		return err
	}

	return nil
}

//go:build !integration
// +build !integration

package lib

import (
	"encoding/xml"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func TestMerge(t *testing.T) {
	oldData := fiber.Map{"message": "hello world"}
	newData := fiber.Map{}

	Merge(oldData, &newData)

	utils.AssertEqual(t, oldData["message"], newData["message"], "Merger")
}

func TestMergeXML(t *testing.T) {
	type s struct {
		XMLName xml.Name
		Message string `xml:"message"`
	}
	oldData := &s{
		Message: "hello World",
	}
	newData := &s{}
	MergeXML(oldData, &newData)

	utils.AssertEqual(t, oldData.Message, newData.Message, "Merge XML")
}

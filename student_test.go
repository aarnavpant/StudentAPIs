package main_test

import (
	"github.com/aarnavpant/StudentAPIs.git/students"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Student", func() {
	var studs *students.Student
	Context("can get student Name", func() {
		BeforeEach(func() {
			studs = &students.Student{
				Name:   "Aarnav Pant",
				RollNo: 1,
			}
		})

		It("can get Name of the Student", func() {
			Expect(studs.GetName()).To(Equal("Aarnav Pant"))
		})

		It("can get RollNo of the student", func() {
			Expect(studs.GetRoll()).To(Equal(1))
		})

		It("can get Student's last Name", func() {
			Expect(studs.GetLastname()).To(Equal("Pant"))
		})

		It("can get Student's last and middle name if present", func() {
			studs = &students.Student{
				Name:   "Aarnav Prasad Pant",
				RollNo: 1,
			}
			Expect(studs.GetLastname()).To(Equal("PrasadPant"))
		})
	})

})

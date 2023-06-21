package src_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/aarnavpant/StudentAPIs.git/src"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Unit", func() {
	var studs *src.Student
	Context("can get student Name", func() {
		BeforeEach(func() {
			studs = &src.Student{
				Name:   "Aarnav Pant",
				RollNo: 1,
			}
		})

		It("can get Name of the Student", func() {
			Expect(studs.GetName()).To(Equal("Aarnav Pant"))
		})

		It("can get Student's last Name", func() {
			Expect(studs.GetLastname()).To(Equal("Pant"))
		})

		It("can get Student's last and middle name if present", func() {
			studs = &src.Student{
				Name:   "Aarnav Prasad Pant",
				RollNo: 1,
			}
			Expect(studs.GetLastname()).To(Equal("PrasadPant"))
		})
	})
	Context("can get student roll no", func() {
		BeforeEach(func() {
			studs = &src.Student{
				Name:   "Aarnav Pant",
				RollNo: 1,
			}
		})
		It("can get RollNo of the student", func() {
			Expect(studs.GetRoll()).To(Equal(1))
		})

	})

	Context("can get response of getAPI", func() {
		Context("can get details of students", func ()  {
			var endpoint string
			var client *http.Client
			BeforeEach(func(){
				muxGet := src.CreateNewMux()
				server := httptest.NewServer(muxGet)
				endpoint = server.URL
				client = server.Client()
			})
			It("can fetch all students", func() {
				resp, err := client.Get(endpoint + "/getStudents")
				Expect(err).NotTo(HaveOccurred())
				defer resp.Body.Close()
	
				// Validate the response
				Expect(resp.StatusCode).To(Equal(http.StatusOK))
			})
			It("can fetch one student", func(){
				
				resp, err := client.Get(endpoint + "/getStudent?id=1")
				Expect(err).NotTo(HaveOccurred())
	
				defer resp.Body.Close() 
				// Validate the response
				Expect(resp.StatusCode).To(Equal(http.StatusOK))
			})
		})
		It("can add student in record", func(){
			muxGet := src.CreateNewMux()
			server := httptest.NewServer(muxGet)
			endpoint := server.URL
			requestBody := []byte(`{"Name": "Dhruv", "RollNo": 5}`)
			client := server.Client()
			resp, err := client.Post(endpoint + "/addStudent","application/json", bytes.NewReader(requestBody))
			Expect(err).NotTo(HaveOccurred())

			defer resp.Body.Close() 
			// Validate the response
			Expect(resp.StatusCode).To(Equal(http.StatusCreated))
		})
		Context("can delete student", func(){
			It("can delete student from record", func(){
				muxGet := src.CreateNewMux()
				server := httptest.NewServer(muxGet)
				endpoint := server.URL
				requestBody := []byte(`{"Name": "Dhruv", "RollNo": 5}`)
				req, error := http.NewRequest("DELETE", endpoint+"/deleteStudent?id=1",bytes.NewReader(requestBody))
				Expect(error).NotTo(HaveOccurred())
				client := server.Client()
				resp, err := client.Do(req)
				Expect(err).NotTo(HaveOccurred())
				defer resp.Body.Close() 
				// Validate the response
				Expect(resp.StatusCode).To(Equal(http.StatusOK))
			})
		})
	})
})

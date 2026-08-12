package main

import "fmt"

type Student struct {
	ID       string
	Name     string
	Grade    float64
	IsActive bool
}

// value receiver buat ambil info
func (s Student) GetInfo() string {
	st := "Tidak Aktif"
	if s.IsActive {
		st = "Aktif"
	}
	return fmt.Sprintf("ID: %s | Nama: %s | Nilai: %.2f | Status: %s", s.ID, s.Name, s.Grade, st)
}

// pointer receiver ubah data
func (s *Student) UpdateGrade(n float64) {
	s.Grade = n
}
func (s *Student) Activate() {
	s.IsActive = true
}
func (s *Student) Deactivate() {
	s.IsActive = false
}

func main() {
	mhs := Student{ID: "5025211001", Name: "Alfin", Grade: 3.50, IsActive: true}
	fmt.Println("Data Awal:", mhs.GetInfo())
	mhs.UpdateGrade(3.85)
	fmt.Println("Setalah Update Grade:", mhs.GetInfo())
	mhs.Deactivate()
	fmt.Println("Setelah Deactivate:", mhs.GetInfo())
	mhs.Activate()
	fmt.Println("Setelah Activate:", mhs.GetInfo())
}

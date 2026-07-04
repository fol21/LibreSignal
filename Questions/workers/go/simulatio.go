package main

import "fmt"

type WorkerRegistration struct {
	w  map[string]*Worker
	ts []int
	// topN []*Worker // can be used as MinHeap for optimal
}

type Worker struct {
	Id        string
	Career    *Career
	Promotion *Role
	Reg       [][2]int
	Worked    int
	Earned    int
}

type Role struct {
	Position     string
	Compensation int
	createdAt    string
}

type Career struct {
	roles []*Role
	hm    map[string]*Role
}

const INV = "invalid_request"
const OKREG = "registered"

func (wr WorkerRegistration) AddWorker(workerId string, position string, compensation int) bool {
	if _, exists := wr.w[workerId]; exists {
		return false
	}
	wr.w[workerId] = &Worker{
		Id: workerId,
		Career: &Career{
			roles: []*Role{},
			hm:    make(map[string]*Role),
		},
		Promotion: &Role{
			Position:     position,
			Compensation: compensation,
			createdAt:    "",
		},
		Reg: [][2]int{},
	}

	return true
}

func (wr WorkerRegistration) Register(workerId string, timestamp int) string {
	if _, exists := wr.w[workerId]; !exists {
		return INV
	}
	p := wr.w[workerId]
	if len(p.Reg) == 0 || p.Reg[len(p.Reg)-1][1] != -1 {
		// Add promotion to career
		p.Career.roles = append(p.Career.roles, p.Promotion)
		p.Promotion = nil

		// add begin of day
		p.Reg = append(p.Reg, [2]int{timestamp, -1})
		return OKREG
	}
	p.Reg[len(p.Reg)-1][1] = timestamp
	p.Worked += timestamp - p.Reg[len(p.Reg)-1][0]
	return OKREG
}

func (wr WorkerRegistration) Get(workerId string) string {
	if _, exists := wr.w[workerId]; !exists {
		return ""
	}
	p := wr.w[workerId]
	return fmt.Sprintf("%d", p.Worked)
}

// func (wr WorkerRegistration) TopNWorkers(n int) []string {

// }

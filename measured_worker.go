package main

import "sync"

type MeasuredWorker struct {
	Worker
	value int
	mutex sync.Mutex
}

func (m *MeasuredWorker) Work() {
	m.Worker.Work()
	m.mutex.Lock()
	m.value++
	m.mutex.Unlock()
}

func (m *MeasuredWorker) Value() int {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m.value
}

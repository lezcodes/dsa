package array_list

import "slices"

import "fmt"

type ArrayList struct {
	data []int
}

func NewArrayList() *ArrayList {
	return &ArrayList{
		data: make([]int, 0),
	}
}

func (al *ArrayList) Add(item int) {
	al.data = append(al.data, item)
}

func (al *ArrayList) Get(index int) (int, error) {
	if index < 0 || index >= len(al.data) {
		return 0, fmt.Errorf("index out of bounds")
	}
	return al.data[index], nil
}

func (al *ArrayList) Set(index int, item int) error {
	if index < 0 || index >= len(al.data) {
		return fmt.Errorf("index out of bounds")
	}
	al.data[index] = item
	return nil
}

func (al *ArrayList) Remove(index int) (int, error) {
	if index < 0 || index >= len(al.data) {
		return 0, fmt.Errorf("index out of bounds")
	}

	item := al.data[index]
	al.data = slices.Delete(al.data, index, index+1)
	return item, nil
}

func (al *ArrayList) Size() int {
	return len(al.data)
}

func (al *ArrayList) IndexOf(item int) int {
	for i, v := range al.data {
		if v == item {
			return i
		}
	}
	return -1
}

func Run() any {
	al := NewArrayList()

	al.Add(10)
	al.Add(20)
	al.Add(30)

	size := al.Size()
	item, _ := al.Get(1)
	al.Set(1, 25)
	index := al.IndexOf(30)
	al.Remove(0)

	return map[string]any{
		"size":        size,
		"item_at_1":   item,
		"index_of_30": index,
		"final_size":  al.Size(),
	}
}

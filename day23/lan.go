package day23

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/makarchuk/aoc2024/pkg/set"
)

type Input struct {
	Connections map[string]set.Set[string]
}

func (i *Input) AddConnection(a, b string) {
	if i.Connections == nil {
		i.Connections = make(map[string]set.Set[string])
	}
	if _, ok := i.Connections[a]; !ok {
		i.Connections[a] = set.New[string]()
	}
	if _, ok := i.Connections[b]; !ok {
		i.Connections[b] = set.New[string]()
	}
	i.Connections[a].Add(b)
	i.Connections[b].Add(a)
}

func ParseInput(in *io.Reader) (*Input, error) {
	i := &Input{}
	scanner := bufio.NewScanner(*in)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			return nil, fmt.Errorf("Invalid input: %s", line)
		}
		a, b := parts[0], parts[1]
		i.AddConnection(a, b)
	}

	return i, nil
}

func Part1(in io.Reader) (string, error) {

	groupsOfThree := 0

	input, err := ParseInput(&in)
	if err != nil {
		return "", err
	}

	for a, connections := range input.Connections {
		connectionsList := connections.List()

		for i, b := range connectionsList {
			for _, c := range connectionsList[i:] {
				if input.Connections[b].Contains(c) {
					if strings.HasPrefix(a, "t") || strings.HasPrefix(b, "t") || strings.HasPrefix(c, "t") {
						groupsOfThree++
					}
				}
			}
		}
	}

	return fmt.Sprintf("%d", groupsOfThree/3), nil
}

func Part2(in io.Reader) (string, error) {
	input, err := ParseInput(&in)
	if err != nil {
		return "", err
	}

	biggest := biggestLanParty(*input)
	members := biggest.Members.List()
	slices.Sort(members)
	return strings.Join(members, ","), nil
}

func biggestLanParty(in Input) Group {
	visited := set.New[string]()
	biggest := &Group{}

	for member, connections := range in.Connections {
		fmt.Printf("Checking %v. Current biggest: %v\n", member, biggest.Size)
		group := &Group{
			Members:           set.From([]string{member}),
			CommonConnections: connections,
			Size:              1,
		}

		biggestChild := group.FindBiggestGroup(in, visited)
		if biggestChild.Size > biggest.Size {
			biggest = biggestChild
		}
	}

	return *biggest
}

func (g *Group) FindBiggestGroup(input Input, visited set.Set[string]) *Group {
	biggest := g
	for member := range g.CommonConnections.Iter() {
		connections := input.Connections[member]
		newMembers := g.Members.Clone()
		newMembers.Add(member)
		newGroup := &Group{
			Members:           newMembers,
			CommonConnections: connections.Intersection(g.CommonConnections),
			Size:              g.Size + 1,
		}
		gid := newGroup.ID()
		if visited.Contains(gid) {
			continue
		}
		visited.Add(gid)
		if newGroup.CommonConnections.Len() == 0 {
			if newGroup.Size > biggest.Size {
				biggest = newGroup
			}
			continue
		}
		biggestChild := newGroup.FindBiggestGroup(input, visited)
		if biggestChild.Size > biggest.Size {
			biggest = biggestChild
		}
	}
	return biggest
}

type Group struct {
	Members           set.Set[string]
	CommonConnections set.Set[string]
	Size              int
}

func (g *Group) Add(member string, connections set.Set[string]) (Group, bool) {
	intersection := g.CommonConnections.Intersection(connections)
	if intersection.Len() == 0 {
		return Group{}, false
	}
	newGroup := Group{}
	newGroup.Members = g.Members.Clone()
	newGroup.Members.Add(member)
	newGroup.CommonConnections = intersection
	newGroup.Size = g.Size + 1
	return newGroup, true
}

func (g *Group) Contains(member string) bool {
	return g.Members.Contains(member)
}

func (g *Group) ID() string {
	members := g.Members.List()
	slices.Sort(members)
	return strings.Join(members, ",")
}

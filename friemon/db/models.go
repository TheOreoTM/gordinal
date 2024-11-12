// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"time"

	"github.com/google/uuid"
)

type Character struct {
	ID               uuid.UUID `json:"id"`
	OwnerID          string    `json:"owner_id"`
	ClaimedTimestamp time.Time `json:"claimed_timestamp"`
	Idx              int32     `json:"idx"`
	CharacterID      int32     `json:"character_id"`
	Level            int32     `json:"level"`
	Xp               int32     `json:"xp"`
	Personality      string    `json:"personality"`
	Shiny            bool      `json:"shiny"`
	IvHp             int32     `json:"iv_hp"`
	IvAtk            int32     `json:"iv_atk"`
	IvDef            int32     `json:"iv_def"`
	IvSpAtk          int32     `json:"iv_sp_atk"`
	IvSpDef          int32     `json:"iv_sp_def"`
	IvSpd            int32     `json:"iv_spd"`
	IvTotal          float64   `json:"iv_total"`
	Nickname         string    `json:"nickname"`
	Favourite        bool      `json:"favourite"`
	HeldItem         int32     `json:"held_item"`
	Moves            []int32   `json:"moves"`
	Color            int32     `json:"color"`
}

type User struct {
	ID            string    `json:"id"`
	Balance       int32     `json:"balance"`
	SelectedID    uuid.UUID `json:"selected_id"`
	OrderBy       int32     `json:"order_by"`
	OrderDesc     bool      `json:"order_desc"`
	ShiniesCaught int32     `json:"shinies_caught"`
	NextIdx       int32     `json:"next_idx"`
}

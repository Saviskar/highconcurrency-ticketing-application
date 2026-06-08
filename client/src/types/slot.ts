export interface Slot {
  ID: number
  TimeSlot: string
  Status: "available" | "booked"
  BookedBy: string
  UpdatedAt: string
}

export interface BookSlotRequest {
  slot_id: number
  user: string
}

export interface BookSlotResponse {
  message: string
  slot: Slot
}

export interface WsMessage {
  event: string
  data: {
    slot_id: number
  }
}

import type { Slot, BookSlotRequest, BookSlotResponse } from "@/types/slot"

const BASE_URL = import.meta.env.VITE_API_URL ?? "http://localhost:8080"

export async function fetchSlots(): Promise<Slot[]> {
  const res = await fetch(`${BASE_URL}/api/v1/slots`)
  if (!res.ok) throw new Error("Failed to fetch slots")
  return res.json()
}

export async function bookSlot(data: BookSlotRequest): Promise<BookSlotResponse> {
  const res = await fetch(`${BASE_URL}/api/v1/slots/book`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  })
  if (res.status === 409) {
    const err = await res.json()
    throw new BookingConflictError(err.error ?? "Slot is currently locked", data.slot_id)
  }
  if (!res.ok) {
    const err = await res.json().catch(() => ({}))
    throw new Error(err.error ?? "Failed to book slot")
  }
  return res.json()
}

export class BookingConflictError extends Error {
  slotId: number
  constructor(message: string, slotId: number) {
    super(message)
    this.name = "BookingConflictError"
    this.slotId = slotId
  }
}

export function getWsUrl(): string {
  const base = BASE_URL.replace(/^http/, "ws")
  return `${base}/api/v1/ws`
}

import { Badge } from "@/components/ui/badge"
import type { Slot } from "@/types/slot"

export function StatusBadge({ slot }: { slot: Slot }) {
  if (slot.Status === "booked") {
    return <Badge variant="destructive">Booked</Badge>
  }
  return <Badge variant="success">Available</Badge>
}

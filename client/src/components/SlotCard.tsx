import { Card, CardContent, CardFooter, CardHeader } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Skeleton } from "@/components/ui/skeleton"
import { StatusBadge } from "@/components/ui/StatusBadge"
import { Clock, User } from "lucide-react"
import type { Slot } from "@/types/slot"

interface SlotCardProps {
  slot: Slot
  onBook: (slot: Slot) => void
  isBooking?: boolean
}

export function SlotCard({ slot, onBook, isBooking }: SlotCardProps) {
  const isAvailable = slot.Status === "available"

  return (
    <Card className={`transition-shadow hover:shadow-md ${!isAvailable ? "opacity-80" : ""}`}>
      <CardHeader className="pb-3">
        <div className="flex items-center justify-between">
          <div className="flex items-center gap-2 text-sm text-muted-foreground">
            <Clock className="h-4 w-4" />
            <span>{slot.TimeSlot}</span>
          </div>
          <StatusBadge slot={slot} />
        </div>
      </CardHeader>

      <CardContent className="pb-3">
        {!isAvailable && (
          <div className="flex items-center gap-2 text-sm">
            <User className="h-4 w-4 text-muted-foreground" />
            <span className="text-muted-foreground">
              Booked by <span className="font-medium text-foreground">{slot.BookedBy}</span>
            </span>
          </div>
        )}
        {isAvailable && (
          <p className="text-sm text-muted-foreground">This slot is open for booking</p>
        )}
      </CardContent>

      <CardFooter>
        <Button
          className="w-full"
          variant={isAvailable ? "default" : "secondary"}
          disabled={!isAvailable || isBooking}
          onClick={() => onBook(slot)}
        >
          {isBooking ? (
            <span className="flex items-center gap-2">
              <span className="h-4 w-4 animate-spin rounded-full border-2 border-current border-t-transparent" />
              Booking...
            </span>
          ) : isAvailable ? (
            "Book Now"
          ) : (
            "Unavailable"
          )}
        </Button>
      </CardFooter>
    </Card>
  )
}

export function SlotCardSkeleton() {
  return (
    <Card>
      <CardHeader className="pb-3">
        <div className="flex items-center justify-between">
          <Skeleton className="h-4 w-32" />
          <Skeleton className="h-5 w-16 rounded-full" />
        </div>
      </CardHeader>
      <CardContent className="pb-3">
        <Skeleton className="h-4 w-40" />
      </CardContent>
      <CardFooter>
        <Skeleton className="h-9 w-full rounded-md" />
      </CardFooter>
    </Card>
  )
}

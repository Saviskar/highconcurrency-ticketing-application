import { SlotCard, SlotCardSkeleton } from "@/components/SlotCard"
import { Button } from "@/components/ui/button"
import { RefreshCw, AlertCircle } from "lucide-react"
import type { Slot } from "@/types/slot"

interface SlotGridProps {
  slots: Slot[] | undefined
  isLoading: boolean
  isError: boolean
  error: Error | null
  bookingSlotId: number | null
  onBook: (slot: Slot) => void
  onRetry: () => void
}

export function SlotGrid({
  slots,
  isLoading,
  isError,
  error,
  bookingSlotId,
  onBook,
  onRetry,
}: SlotGridProps) {
  if (isLoading) {
    return (
      <div className="grid gap-4 sm:grid-cols-2">
        <SlotCardSkeleton />
        <SlotCardSkeleton />
        <SlotCardSkeleton />
        <SlotCardSkeleton />
      </div>
    )
  }

  if (isError) {
    return (
      <div className="flex flex-col items-center gap-4 rounded-lg border border-destructive/50 bg-destructive/5 px-6 py-12 text-center">
        <AlertCircle className="h-10 w-10 text-destructive" />
        <div>
          <h3 className="font-semibold">Failed to load slots</h3>
          <p className="mt-1 text-sm text-muted-foreground">
            {error?.message ?? "Something went wrong"}
          </p>
        </div>
        <Button variant="outline" onClick={onRetry}>
          <RefreshCw className="mr-2 h-4 w-4" />
          Try Again
        </Button>
      </div>
    )
  }

  if (!slots || slots.length === 0) {
    return (
      <div className="rounded-lg border px-6 py-12 text-center">
        <h3 className="font-semibold">No slots available</h3>
        <p className="mt-1 text-sm text-muted-foreground">
          Check back later for available time slots.
        </p>
      </div>
    )
  }

  return (
    <div className="grid gap-4 sm:grid-cols-2">
      {slots.map((slot) => (
        <SlotCard
          key={slot.ID}
          slot={slot}
          onBook={onBook}
          isBooking={bookingSlotId === slot.ID}
        />
      ))}
    </div>
  )
}

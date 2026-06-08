import { useState, useCallback } from "react"
import { useSlots } from "@/hooks/useSlots"
import { useBookSlot, BookingConflictError } from "@/hooks/useBookSlot"
import { useWebSocket } from "@/hooks/useWebSocket"
import { useUserStore } from "@/store/userStore"
import { Header } from "@/components/Header"
import { SlotGrid } from "@/components/SlotGrid"
import { BookingDialog } from "@/components/BookingDialog"
import { Toaster } from "@/components/ui/sonner"
import { toast } from "sonner"
import type { Slot } from "@/types/slot"

export default function App() {
  const { connected } = useWebSocket()
  const { data: slots, isLoading, isError, error, refetch } = useSlots()
  const bookMutation = useBookSlot()
  const { name: savedName, setName } = useUserStore()

  const [selectedSlot, setSelectedSlot] = useState<Slot | null>(null)

  const handleBook = useCallback(
    (slot: Slot) => {
      setSelectedSlot(slot)
    },
    [],
  )

  const handleConfirm = useCallback(
    async (slot: Slot, name: string) => {
      if (!name.trim()) return

      try {
        const result = await bookMutation.mutateAsync({
          slot_id: slot.ID,
          user: name,
        })
        toast.success(result.message ?? "Slot booked successfully!")
        setSelectedSlot(null)
      } catch (err) {
        if (err instanceof BookingConflictError) {
          toast.error("Someone else is booking this slot. Please try again.")
        } else {
          toast.error(err instanceof Error ? err.message : "Failed to book slot")
        }
      }
    },
    [bookMutation],
  )

  const handleClose = useCallback(() => setSelectedSlot(null), [])
  const handleChangeName = useCallback(
    (name: string) => setName(name),
    [setName],
  )
  const handleRetry = useCallback(() => refetch(), [refetch])

  return (
    <div className="min-h-screen bg-background">
      <Header wsConnected={connected} />

      <main className="mx-auto max-w-3xl px-4 py-8">
        <div className="mb-6">
          <h2 className="text-2xl font-bold tracking-tight">Available Slots</h2>
          <p className="mt-1 text-sm text-muted-foreground">
            Select a time slot to book your appointment.
          </p>
        </div>

        <SlotGrid
          slots={slots}
          isLoading={isLoading}
          isError={isError}
          error={error}
          bookingSlotId={bookMutation.isPending ? bookMutation.variables?.slot_id ?? null : null}
          onBook={handleBook}
          onRetry={handleRetry}
        />
      </main>

      <BookingDialog
        slot={selectedSlot}
        userName={savedName}
        isPending={bookMutation.isPending}
        onConfirm={handleConfirm}
        onClose={handleClose}
        onChangeName={handleChangeName}
      />

      <Toaster richColors position="top-right" />
    </div>
  )
}

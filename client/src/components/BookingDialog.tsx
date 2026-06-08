import { useState, useEffect } from "react"
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogDescription,
  DialogFooter,
} from "@/components/ui/dialog"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Clock, User } from "lucide-react"
import type { Slot } from "@/types/slot"

interface BookingDialogProps {
  slot: Slot | null
  userName: string
  isPending: boolean
  onConfirm: (slot: Slot, name: string) => void
  onClose: () => void
  onChangeName: (name: string) => void
}

export function BookingDialog({
  slot,
  userName,
  isPending,
  onConfirm,
  onClose,
  onChangeName,
}: BookingDialogProps) {
  const [name, setName] = useState(userName)
  const open = slot !== null

  useEffect(() => {
    if (open) setName(userName)
  }, [open, userName])

  if (!slot) return null

  return (
    <Dialog open={open} onOpenChange={(o) => !o && !isPending && onClose()}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Confirm Booking</DialogTitle>
          <DialogDescription>
            You are about to book this time slot.
          </DialogDescription>
        </DialogHeader>

        <div className="space-y-4 py-2">
          <div className="flex items-center gap-2 rounded-md bg-muted px-3 py-2 text-sm">
            <Clock className="h-4 w-4 text-muted-foreground" />
            <span className="font-medium">{slot.TimeSlot}</span>
          </div>

          <div className="space-y-1.5">
            <label className="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70">
              Your name
            </label>
            <div className="relative">
              <User className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
              <Input
                className="pl-9"
                placeholder="Enter your name"
                value={name}
                onChange={(e) => {
                  setName(e.target.value)
                  onChangeName(e.target.value)
                }}
                disabled={isPending}
              />
            </div>
          </div>
        </div>

        <DialogFooter>
          <Button variant="outline" onClick={onClose} disabled={isPending}>
            Cancel
          </Button>
          <Button
            onClick={() => onConfirm(slot, name)}
            disabled={!name.trim() || isPending}
          >
            {isPending ? (
              <span className="flex items-center gap-2">
                <span className="h-4 w-4 animate-spin rounded-full border-2 border-current border-t-transparent" />
                Booking...
              </span>
            ) : (
              "Confirm Booking"
            )}
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}

import { useMutation, useQueryClient } from "@tanstack/react-query"
import { bookSlot, BookingConflictError } from "@/lib/api"
import type { BookSlotRequest } from "@/types/slot"

export function useBookSlot() {
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: (data: BookSlotRequest) => bookSlot(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["slots"] })
    },
    onError: () => {
      queryClient.invalidateQueries({ queryKey: ["slots"] })
    },
    retry: false,
  })
}

export { BookingConflictError }

import { useQuery } from "@tanstack/react-query"
import { fetchSlots } from "@/lib/api"

export function useSlots() {
  return useQuery({
    queryKey: ["slots"],
    queryFn: fetchSlots,
    refetchInterval: 30_000,
  })
}

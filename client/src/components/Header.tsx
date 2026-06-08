import { useState } from "react"
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { useUserStore } from "@/store/userStore"
import { ConnectionStatus } from "@/components/ConnectionStatus"
import { Calendar } from "lucide-react"

interface HeaderProps {
  wsConnected: boolean
}

export function Header({ wsConnected }: HeaderProps) {
  const { name, setName } = useUserStore()
  const [localName, setLocalName] = useState(name)

  return (
    <header className="border-b">
      <div className="mx-auto flex max-w-5xl items-center justify-between px-4 py-3">
        <div className="flex items-center gap-2">
          <Calendar className="h-5 w-5 text-primary" />
          <h1 className="text-lg font-semibold">Booking System</h1>
        </div>

        <div className="flex items-center gap-4">
          <ConnectionStatus connected={wsConnected} />

          <div className="flex items-center gap-2">
            <Input
              placeholder="Your name"
              value={localName}
              onChange={(e) => setLocalName(e.target.value)}
              className="h-8 w-36 text-sm"
            />
            <Button
              size="sm"
              variant="outline"
              onClick={() => setName(localName)}
              disabled={!localName.trim()}
            >
              Set
            </Button>
          </div>
        </div>
      </div>
    </header>
  )
}

"use client"
import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select"
import { University } from "./universityCard"
import RankInput from "./rankInput"
import UniversityCardsContainer from "./universityCardsContainer"
import { Separator } from "../ui/separator"
import { useState } from "react"
import { UniqueIdentifier } from "@dnd-kit/core"
import { convertCapitalCamelCaseToSnakeCase } from "@/utils/commons"
import { useRouter } from "next/navigation"

interface SidebarProps {
  countries: string[] | null
  fields: string[]
  universities: University[] | null
  handleMapView: (longitude: number, latitude: number, zoom?: number) => void;
}

export function Sidebar({ countries, fields, universities, handleMapView }: SidebarProps) {
  const [universityNameInput, setUniversityNameInput] = useState("")
  const [regionInput, setRegionInput] = useState("")
  const [rankInput, setRankInput] = useState<UniqueIdentifier[]>(fields)
  const urlParams = new URLSearchParams()
  const router = useRouter()

  const handleSearch = () => {
    const rankInputStr = convertCapitalCamelCaseToSnakeCase(rankInput.join(":asc,") + ":asc")
    console.log(rankInputStr)
    urlParams.set('sort', rankInputStr)
    urlParams.set('country', regionInput)

    router.push(`?${urlParams.toString()}`)
  }

  return (
    <Card className="w-full h-full flex flex-col">
      <CardHeader>
        <CardTitle className="text-2xl">University Mapper</CardTitle>
        <CardDescription>Find the best university for you based on your location</CardDescription>
      </CardHeader>
      <CardContent className="flex-1 flex flex-col min-h-0 overflow-hidden">
        <div className="flex flex-col gap-4 mt-1">
          <div className="flex gap-4">
            <Input id="university_name" placeholder="Search Universities..."
              value={universityNameInput} onChange={(e) => setUniversityNameInput(e.target.value)} />
            <Button onClick={() => {
              handleSearch()
            }}>Search</Button>
          </div>
          <Select onValueChange={setRegionInput}>
            <SelectTrigger id="region">
              <SelectValue placeholder="Select Region" />
            </SelectTrigger>
            <SelectContent position="popper">
              {countries?.map(country => (
                <SelectItem key={country} value={country}>{country}</SelectItem>
              ))}
            </SelectContent>
          </Select>
          <div className="w-full">
            <RankInput items={rankInput} setItems={setRankInput} />
          </div>
        </div>
        <Separator className="my-4  h-[2px]" />
        <UniversityCardsContainer universities={universities} handleMapView={handleMapView} />
      </CardContent>
    </Card>
  )
}



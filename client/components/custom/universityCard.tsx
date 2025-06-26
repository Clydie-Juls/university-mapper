import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Label } from "../ui/label"
import { MapPin, Pin } from "lucide-react";

export type University = {
  name: string;
  overallRank: number;
  country: string
  ranks: Record<string, number>;
  latitude: string;
  longitude: string;
};

type UniversityCardProps = {
  university: University;
  handleMapView: (longitude: number, latitude: number, zoom?: number) => void;
}


export default function UniversityCard({ university, handleMapView }: UniversityCardProps) {
  return (
    <Card className="w-full cursor-pointer hover:brightness-90 transition-all" onClick={() => {
      handleMapView(Number(university.longitude), Number(university.latitude), 15)
    }}>
      <CardHeader>
        <CardTitle>{university.name}</CardTitle>
        <CardDescription>Rank #{university.overallRank === 0 || university.overallRank === 999999 ? "unknown" : university.overallRank}</CardDescription>
        <CardDescription className="flex items-center gap-2">
          <Pin size={14} color="#f23636" />
          ({university.latitude}, {university.longitude})
        </CardDescription>
        <CardDescription className="flex items-center gap-2">
          <MapPin size={14} color="#f23636" />
          {university.country}
        </CardDescription>
      </CardHeader>
      <CardContent>
        <div className="flex flex-col gap-2 justify-center">
          {Object.entries(university.ranks).map(([subject, rank]) => (
            <div key={subject} className="flex justify-between items-center">
              <Label>{subject}:</Label>
              <div className=" bg-violet-200 px-2 h-6 rounded-full flex items-center justify-center">
                <Label>{rank === 0 || rank === 999999 ? "unknown" : "#" + rank}</Label>
              </div>
            </div>
          ))}
        </div>
      </CardContent>
    </Card>
  )
}


import { Feature } from "ol"
import { Point } from "ol/geom"

import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Button } from "../ui/button"
import Link from "next/link"

interface UniversityPopupCardProps {
  feature: Feature<Point>
}

export default function UniversityPopupCard({ feature }: UniversityPopupCardProps) {
  const rank = feature.get('overall-rank') == '0' ||
    feature.get('overall-rank') == '999999' ? "unkown" :
    "#" + feature.get('overall-rank')

  const createStreetViewUrl = (lat: string, lng: string) => 
  `https://www.google.com/maps/@${lat},${lng},3a,75y,90t/data=!3m6!1e1!3m3!1sAF1QipMk6RmDcfGcXswvlKTzGqR6Mg2dtFbpIDEX_z5y!2e10!3e12`;
  return (
    <Card className="w-[280px] opacity-80 hover:opacity-100 duration-200">
      <CardHeader>
        <CardTitle>{feature.get('name')}</CardTitle>
        <CardDescription>rank: {rank}</CardDescription>
      </CardHeader>
      <CardContent>
        <Button asChild>
          <Link target="_blank" href={createStreetViewUrl(feature.get('lat'), feature.get('lng'))}>Check in Google Maps</Link>
        </Button>
      </CardContent>
    </Card>
  )
}


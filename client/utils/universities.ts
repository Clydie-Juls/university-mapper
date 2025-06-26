import { UniversityJSON } from "@/app/page";
import { University } from "@/components/custom/universityCard";

export function transformUniversityData(university: UniversityJSON): University {
  const result: University = {
    name: university.Name,
    overallRank: university.OverallRank,
    country: university.Country,
    latitude: university.Latitude,
    longitude: university.Longitude,
    ranks: {}
  }

  for (const key in university) {
    if (key.toLowerCase().includes("rank") && !key.toLowerCase().includes("overallrank")) {
      let newKey = key.replace("Rank", "")
      newKey = newKey.replace(/([a-z])([A-Z])/g, '$1 $2');
      result.ranks[newKey] = university[key]
    }
  }

  return result
}

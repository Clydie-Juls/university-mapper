import { UniversityJSON } from "@/app/page"
import { University } from "@/components/custom/universityCard"
import { transformUniversityData } from "@/utils/universities"

function getSortQueryParam(university: UniversityJSON, urlParams: string): string {
  if (!urlParams) {
    let query: string[] = ["overall_rank:asc"]
    for (const key in university) {
      if (key == "OverallRank") {
        continue
      }
      const newKey = key.replace(/([A-Z])/g, (match) => `_${match.toLowerCase()}`).slice(1)
      query.push(`${newKey}:asc`)
    }

    return `sort=${query.join(",").toLowerCase()}`
  }

  return urlParams
}

export async function getCountries(): Promise<string[]> {
  let data = await fetch('http://localhost:8080/api/v1/universities/countries', {
    cache: 'no-store'
  })
  if (!data.ok) throw new Error("Unable to fetch countries")
  return await data.json()
}

export async function getUniversityFields(): Promise<string[]> {
  const data = await fetch('http://localhost:8080/api/v1/universities/fields', {
    cache: 'no-store'
  })
  if (!data.ok) throw new Error("Unable to fetch countries")
  return await data.json()
}

export async function getUniversities(urlParams: string): Promise<University[]> {
  const sortQueryParam = getSortQueryParam({} as UniversityJSON, urlParams)
  const data = await fetch(`http://localhost:8080/api/v1/universities/?${sortQueryParam}`, {
    cache: 'no-store'
  })
  const universitiesJson: UniversityJSON[] = await data.json()
  let universities: University[] = [];
  universitiesJson.forEach((uniJson) => {
    universities.push(transformUniversityData(uniJson))
  })
  return universities
}

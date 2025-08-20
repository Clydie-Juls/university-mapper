import UniversityMap from "@/components/custom/universityMap";
import { getCountries, getUniversities, getUniversityFields } from "@/lib/api";

export interface UniversityJSON {
  Name: string;
  OverallRank: number;
  Latitude: string;
  Longitude: string;
  [key: string]: any;
}

interface HomeSearchParams {
  country: string
  sort: string
}

export default async function Home({ searchParams }: { searchParams: Promise<HomeSearchParams> }) {
  const resolvedSearchParams = await searchParams;
  const filteredParams = Object.fromEntries(
    Object.entries(resolvedSearchParams).filter(([_, value]) => value !== undefined) as [string, string][]
  );

  // Convert to URLSearchParams
  const urlParams = new URLSearchParams(filteredParams).toString();
  console.log(urlParams)

  const countries = await getCountries()
  console.log(countries)
  const fields = await getUniversityFields()
  console.log(fields)
  const universities = await getUniversities(urlParams)
  console.log(universities)


  return (
    <div className=" w-screen h-screen bg-black p-2 grid grid-cols-[1fr_400px] gap-2 overflow-hidden box-border">
      <UniversityMap countries={countries} fields={fields} universities={universities} />
    </div>
  )
}

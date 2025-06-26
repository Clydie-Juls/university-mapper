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

export default async function Home({ searchParams }: {searchParams: HomeSearchParams}) {
  const filteredParams = Object.fromEntries(
    Object.entries(searchParams).filter(([_, value]) => value !== undefined) as [string, string][]
  );

  // Convert to URLSearchParams
  const urlParams = new URLSearchParams(filteredParams).toString();
  console.log(urlParams)

  const countries = await getCountries()
  const fields = await getUniversityFields()
  const universities = await getUniversities(urlParams)


  return (
    <div className=" w-screen h-screen bg-black p-2 grid grid-cols-[1fr_400px] gap-2 overflow-hidden box-border">
      <UniversityMap countries={countries} fields={fields} universities={universities}/>
    </div>
  )
}

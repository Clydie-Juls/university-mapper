"use client"
import { useEffect, useState } from "react"
import UniversityCard, { University } from "./universityCard"
import { useInView } from "react-intersection-observer";
import Spinner from "./spinner"
import { Input } from "../ui/input";
import { useDebouncedCallback } from 'use-debounce';
import Fuse from "fuse.js";

interface UniversityCardsContainerProps {
  universities: University[] | null
  handleMapView: (longitude: number, latitude: number, zoom?: number) => void;
}

const noOfVisibleElems = 20

function isQueriedAll(searchTerm: string, visibleCardNo: number, searchedVisibleCardsNo: number, uniLength: number, queryLength: number): boolean {
  if (searchTerm != "") {
    return searchedVisibleCardsNo >= queryLength
  }

  return visibleCardNo >= uniLength
}

export default function UniversityCardsContainer({
  universities, handleMapView }: UniversityCardsContainerProps) {
  universities = universities || []
  const [visibleCardsNo, setVisibleCardsNo] = useState<number>(Math.min(universities.length, noOfVisibleElems))
  const [searchedVisibleCardsNo, setSearchedVisibleCardsNo] = useState<number>(Math.min(universities.length, noOfVisibleElems))
  const [searchTerm, setSearchTerm] = useState<string>("")
  const [query, setQuery] = useState<University[]>([])

  const { ref, inView } = useInView()

  const fuse = new Fuse(universities, {
    keys: ['name']
  })

  const onHandleSearch = useDebouncedCallback((term: string) => {
    const unis = fuse.search(term).map(res => res.item)
    console.log(unis)
    setQuery(unis)
    setSearchTerm(term)
    setSearchedVisibleCardsNo(Math.min(universities.length, noOfVisibleElems))
  }, 200)

  useEffect(() => {
    if (searchTerm == "") {
      setVisibleCardsNo(Math.min(noOfVisibleElems + visibleCardsNo, universities.length))
    } else {
      setSearchedVisibleCardsNo(Math.min(noOfVisibleElems + searchedVisibleCardsNo, query.length))
    }
  }, [inView])

  return (
    <>
      <Input type="text" onChange={(e) => onHandleSearch(e.target.value)} placeholder="Filter Universities" />
      <div className="flex-1 min-h-0 flex flex-col gap-4 my-4 overflow-y-auto overflow-x-hidden">
        {searchTerm == "" ?
          universities?.slice(0, visibleCardsNo).map(university => (
            <UniversityCard handleMapView={handleMapView} key={university.name} university={university} />
          )) :
          query?.slice(0, searchedVisibleCardsNo).map(university => (
            <UniversityCard handleMapView={handleMapView} key={university.name} university={university} />
          ))
        }
        <div ref={ref} className="w-full flex items-center justify-center" style={{
          display: isQueriedAll(searchTerm, visibleCardsNo, searchedVisibleCardsNo, universities.length, query.length) ? 'none' : 'flex'
        }}>
          <Spinner />
        </div>
      </div>
    </>
  )
}


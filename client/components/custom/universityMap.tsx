"use client"
import React, { useRef } from 'react'
import OpenLayersMap from './map'
import { Sidebar } from './sidebar'
import { University } from './universityCard'
import { Map } from 'ol'
import { fromLonLat } from 'ol/proj'

interface UniversityMapProps {
  universities: University[],
  countries: string[] | null,
  fields: string[],
}

export default function UniversityMap({ universities, countries, fields }: UniversityMapProps) {
  const mapRef = useRef<Map>(null)
  const handleFlyTo = (longitude: number, latitude: number, zoom = 10) => {
    if (mapRef.current) {
      const point = fromLonLat([longitude, latitude])
      const view = mapRef.current.getView();
      view.animate({
        center: point,
        zoom: zoom,
        duration: 1000,
      });
    }
  };
  return (
    <>
      <div className="w-full h-full rounded-xl overflow-hidden">
        <OpenLayersMap universities={universities} mapRefShare={mapRef} />
      </div>
      <div className="w-full h-screen">
        <Sidebar countries={countries} fields={fields} universities={universities} handleMapView={handleFlyTo} />
      </div>
    </>
  )
}


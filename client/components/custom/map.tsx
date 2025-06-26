'use client'
import React, { useEffect, useRef, useState } from 'react';
import { createRoot } from "react-dom/client";
import 'ol/ol.css';
import { Feature, Map, Overlay, View } from 'ol';
import TileLayer from 'ol/layer/Tile';
import OSM from 'ol/source/OSM';
import { University } from './universityCard';
import VectorSource from 'ol/source/Vector';
import { Point } from 'ol/geom';
import VectorLayer from 'ol/layer/Vector';
import { fromLonLat } from 'ol/proj';
import UniversityPopupCard from './universityPopupCard';
import {useSearchParams } from 'next/navigation';
import { Coordinate } from 'ol/coordinate';

interface OPMapProps {
  universities: University[],
  mapRefShare: React.RefObject<Map | null>
}

function OPMap({ universities, mapRefShare }: OPMapProps) {
  const mapRef = useRef<HTMLDivElement>(null);
  const urlParams = useSearchParams()
  const [center, setCenter] = useState<Coordinate>([34.08, 34.08])
  const [zoom, setZoom] = useState<number>(5)
  const [isSetNavigation, setIsSetNavigation] = useState<Boolean>(false)

  useEffect(() => {
    let currCenter = center
    console.log(navigator.geolocation, isSetNavigation)

    console.log(currCenter)
    const view = new View({
      center: currCenter,
      zoom: zoom,
    })

    if (navigator.geolocation && !isSetNavigation) {
      navigator.geolocation.getCurrentPosition((position) => {
        const {latitude, longitude} = position.coords
        currCenter = fromLonLat([longitude, latitude])
        setCenter(currCenter)
        view.setCenter(currCenter)
        setIsSetNavigation(true)
        console.log("GRRRR")
      })
    }

    const map = new Map({
      target: mapRef.current as HTMLDivElement,
      layers: [
        new TileLayer({
          source: new OSM({
            url: 'https://{a-c}.tile.openstreetmap.org/{z}/{x}/{y}.png?' + new Date().getTime(),
          }),
        }),
      ],
      view: view,
    });

    const updateMapView = () => {
      const view = map.getView();
      const currentCenter = view.getCenter();  // Gets the current center [longitude, latitude]
      const currentZoom = view.getZoom(); // Gets the current zoom level
      if (currentCenter) setCenter(currentCenter);
      if (currentZoom) setZoom(currentZoom);
    };

    // Trigger the update when the view changes
    map.on('moveend', updateMapView);

    const vectorSource = new VectorSource()

    const createPopupElement = (feature: Feature<Point>) => {
      const div = document.createElement("div");
      const root = createRoot(div); // React 18+ (for older versions, use ReactDOM.render)
      root.render(<UniversityPopupCard feature={feature} />);

      return div;
    };

    universities?.forEach((university) => {
      const lat = university.latitude
      const lng = university.longitude
      if (lat != "" && lng != "") {
        const point = [Number(lng), Number(lat)]
        const feature = new Feature({
          geometry: new Point(fromLonLat(point)),
          name: university.name,
          "overall-rank": university.overallRank,
          lat: university.latitude,
          lng: university.longitude,
        })

        const popup = new Overlay({
          element: createPopupElement(feature),
          position: fromLonLat(point),
          autoPan: false
        })

        const elem = popup.getElement()

        if (elem) {
          const parentElem = elem.parentElement
          if (parentElem) {
            parentElem.className = 'z-10 hover:z-50'
          }
        }


        popup.getElement()?.addEventListener('click', () => {
          map.getView().animate({
            center: fromLonLat(point),
            zoom: 20,
            duration: 1000
          })
        })


        map.addOverlay(popup)

        vectorSource.addFeature(feature)
        mapRefShare.current = map
      }
    })

    const vectorLayer = new VectorLayer({
      source: vectorSource,
    });

    map.addLayer(vectorLayer)

    const zoomThreshold = 8

    const updatePopupVisibility = () => {
      const extent = map.getView().calculateExtent(map.getSize());
      const zoom = view.getZoom();

      map.getOverlays().forEach((popup) => {

        const overlay = popup
        const coordinate = popup.getPosition()
        const element = overlay.getElement();
        if (element && zoom && zoom >= zoomThreshold) {
          if (coordinate && coordinate[0] >= extent[0] && coordinate[0] <= extent[2] &&
            coordinate[1] >= extent[1] && coordinate[1] <= extent[3]) {
            element.style.display = 'block'; // Show popup
          } else {
            element.style.display = 'none'; // Hide popup
          }
        } else if (element) {
          element.style.display = 'none'; // Hide popup
        }
      });
    };

    // Listen to map view changes to update visibility
    map.getView().on('change:center', updatePopupVisibility);
    map.getView().on('change:resolution', updatePopupVisibility);

    updatePopupVisibility()


    return () => map.setTarget(undefined);
  }, [urlParams.toString()]);

  return <div ref={mapRef} className='map h-full rounded-xl' id='map' />;
}

export default OPMap

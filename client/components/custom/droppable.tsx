"use client"
import { useDroppable } from '@dnd-kit/core'
import React from 'react'

interface DroppableProps {
  children: React.ReactNode; // `children` is of type React.ReactNode
  id: string
}

export default function Droppable(props: DroppableProps) {
  const {isOver, setNodeRef} = useDroppable({
    id: props.id
  })

  const style = {
    color: isOver ? 'green' : 'black'
  }

  return (
    <div ref={setNodeRef} style={style}>{props.children}</div>
  )
}


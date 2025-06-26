import { useDraggable } from '@dnd-kit/core'
import React from 'react'

interface DraggableProps {
  children: React.ReactNode
}

export default function Draggable(props: DraggableProps) {
  const { attributes, listeners, setNodeRef, transform } = useDraggable({
    id: "draggable"
  })

  const style = transform ? {
    transform: `translate3d(${transform?.x}px, ${transform?.y}px, 0)`
  } : {}

  return (
    <div ref={setNodeRef} style={style} {...listeners} {...attributes}>{props.children}</div>
  )
}


'use client'

import React, { Dispatch, SetStateAction, useState } from 'react'
import { closestCenter, DndContext, DragEndEvent, KeyboardSensor, PointerSensor, UniqueIdentifier, useSensor, useSensors } from '@dnd-kit/core'
import { Button } from "@/components/ui/button"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { arrayMove, SortableContext, sortableKeyboardCoordinates, verticalListSortingStrategy } from '@dnd-kit/sortable';
import SortableItem from './sortableItem'
import { Label } from '../ui/label'
import { ChevronDown } from 'lucide-react'

interface RankInputProps {
  items: UniqueIdentifier[]
  setItems: Dispatch<SetStateAction<UniqueIdentifier[]>>
}

export default function RankInput({items, setItems}: RankInputProps) {
  const sensors = useSensors(
    useSensor(PointerSensor),
    useSensor(KeyboardSensor, {
      coordinateGetter: sortableKeyboardCoordinates
    })
  )

  function handleDragEnd(event: DragEndEvent) {
    const { active, over } = event;

    if (active.id !== over?.id) {
      setItems((items) => {
        const oldIndex = items?.indexOf(active.id);
        const newIndex = items?.indexOf(over?.id || '');
        return arrayMove(items, oldIndex, newIndex);
      });
    }
  }

  return (
    <DndContext onDragEnd={handleDragEnd} sensors={sensors} collisionDetection={closestCenter}>

      <SortableContext items={items || []}
        strategy={verticalListSortingStrategy}
      >
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant="outline" className='w-full flex justify-between'>
              <div>Sort University Ranks/Names/etc.</div>
              <ChevronDown color='grey'/>
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent className="w-80">
            <DropdownMenuLabel className='text-center'>Sort by Priority</DropdownMenuLabel>
            <DropdownMenuSeparator />
            <DropdownMenuGroup>

              {items?.map(id => <SortableItem key={id} id={id}>
                <DropdownMenuItem>
                  <div key={id} className='w-full flex justify-between p-1'>
                    <Label>{id}</Label>
                    <Label>≡</Label>
                  </div>
                </DropdownMenuItem>
              </SortableItem>)}
            </DropdownMenuGroup>
          </DropdownMenuContent>
        </DropdownMenu>
      </SortableContext>
    </DndContext>
  )
}


'use client';

import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import { Compiler } from '@/types/api/compiler';
import { Language } from '@/types/api/language';
export function CompilerSelect({
  lang,
  compiler,
  setCompiler,
}: {
  lang: Language;
  compiler: Compiler;
  setCompiler(l: Compiler): void;
}) {
  return (
    <Select
      onValueChange={v => {
        setCompiler(
          lang.compilers.find(c => v === c.docker_image_name) ??
            lang.compilers[0],
        );
      }}
      value={compiler.docker_image_name}
    >
      <SelectTrigger className='w-[180px] bg-white!'>
        <SelectValue />
      </SelectTrigger>
      <SelectContent>
        <SelectGroup>
          <SelectLabel>Компиляторы</SelectLabel>
          {lang.compilers.map(c => (
            <SelectItem value={c.docker_image_name} key={c.id}>
              {c.docker_image_name}
            </SelectItem>
          ))}
        </SelectGroup>
      </SelectContent>
    </Select>
  );
}

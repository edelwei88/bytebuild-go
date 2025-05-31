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
export function LanguageSelect({
  langs,
  language,
  setLanguage,
  setCompiler,
}: {
  langs: Language[];
  language: Language;
  setLanguage(l: Language): void;
  setCompiler(c: Compiler): void;
}) {
  return (
    <Select
      onValueChange={v => {
        setLanguage(langs.find(l => v === l.name) ?? langs[0]);
        setCompiler(
          langs.find(l => v === l.name)?.compilers[0] ?? langs[0].compilers[0],
        );
      }}
      value={language.name}
    >
      <SelectTrigger className='w-[180px] bg-white!'>
        <SelectValue />
      </SelectTrigger>
      <SelectContent>
        <SelectGroup>
          <SelectLabel>Языки</SelectLabel>
          {langs.map(l => (
            <SelectItem value={l.name} key={l.id}>
              {l.name}
            </SelectItem>
          ))}
        </SelectGroup>
      </SelectContent>
    </Select>
  );
}

import { Compiler } from './compiler';

export interface Language {
  id: number;
  name: string;
  file_extension: string;
  compilers: Compiler[];
  monaco_name: string;
}

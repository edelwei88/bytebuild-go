import { Compiler } from './compiler';

export interface Compile {
  id: number;
  compiler: Compiler;
  args: string;
  exit_code: number;
  source_code: string;
  stdout: string;
  stderr: string;
}

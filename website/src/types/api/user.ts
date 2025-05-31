import { Compile } from './compile';
import { Role } from './role';

export interface User {
  id: number;
  username: string;
  email: string;
  role: Role;
  compiles: Compile[];
}

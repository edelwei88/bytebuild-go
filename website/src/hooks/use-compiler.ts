import { useState, useRef } from 'react';
import { Compiler as CompilerType } from '@/types/api/compiler';
import { Language } from '@/types/api/language';
import { toast } from 'sonner';

interface UseCompilerProps {
  languages: Language[];
  sourceCodeRef?: string;
  stdoutRef?: string;
  stderrRef?: string;
  argsRef?: string;
  compilerRef?: CompilerType;
  languageRef?: Language;
}

export function useCompiler({
  languages,
  sourceCodeRef,
  stdoutRef,
  stderrRef,
  argsRef,
  compilerRef,
  languageRef,
}: UseCompilerProps) {
  const [sourceCode, setSourceCode] = useState(
    sourceCodeRef ?? 'console.log("Hello, world!")',
  );
  const [language, setLanguage] = useState<Language>(
    languageRef ?? languages[0],
  );
  const [compiler, setCompiler] = useState<CompilerType>(
    compilerRef ?? languages[0].compilers[0],
  );
  const [args, setArgs] = useState(argsRef ?? '');
  const [stdout, setStdout] = useState(stdoutRef ?? '');
  const [stderr, setStderr] = useState(stderrRef ?? '');
  const editorRef = useRef(null);
  const [isCompiling, setIsCompiling] = useState(false);

  const handleEditorDidMount = (editor: any) => {
    editorRef.current = editor;
    editor.focus();
    editor.RenderMinimap = 0;
  };

  const handleOnChange = (value: string | undefined) => {
    if (value) setSourceCode(value);
  };

  const compile = async () => {
    if (isCompiling) return;
    
    try {
      setIsCompiling(true);
      const res = await fetch('http://localhost:3001/user/compile', {
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include',
        method: 'post',
        body: JSON.stringify({
          language: language.name,
          compiler: compiler.docker_image_name,
          source_code: sourceCode,
          args: args,
        }),
      });
      
      if (!res.ok) {
        toast.error('Ошибка компиляции', {
          description: 'Ошибка компиляции на сервере'
        });
        return;
      }
      
      const json = await res.json();
      setStdout(json.stdout);
      setStderr(json.stderr);
      toast.success('Успешная компиляция', {
        duration: 1000
      });
    } catch (e) {
      toast.error('Неожиданная ошибка', {
        description: String(e)
      });
      console.error(e);
    } finally {
      setIsCompiling(false);
    }
  };

  return {
    sourceCode,
    setSourceCode,
    language,
    setLanguage,
    compiler,
    setCompiler,
    args,
    setArgs,
    stdout,
    stderr,
    editorRef,
    isCompiling,
    handleEditorDidMount,
    handleOnChange,
    compile
  };
} 
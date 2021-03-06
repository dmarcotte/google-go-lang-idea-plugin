package main
type T interface {
	Read(b Buffer) bool
	Write(b Buffer) bool
	Close()
}
/**-----
Go file
  PackageDeclaration(main)
    PsiElement(KEYWORD_PACKAGE)('package')
    PsiWhiteSpace(' ')
    PsiElement(IDENTIFIER)('main')
  PsiWhiteSpace('\n')
  TypeDeclarationsImpl
    PsiElement(KEYWORD_TYPE)('type')
    PsiWhiteSpace(' ')
    TypeSpecImpl
      TypeNameDeclaration(T)
        PsiElement(IDENTIFIER)('T')
      PsiWhiteSpace(' ')
      TypeInterfaceImpl
        PsiElement(KEYWORD_INTERFACE)('interface')
        PsiWhiteSpace(' ')
        PsiElement({)('{')
        PsiWhiteSpace('\n')
        PsiWhiteSpace('\t')
        FunctionDeclaration(Read)
          LiteralIdentifierImpl
            PsiElement(IDENTIFIER)('Read')
          PsiElement(()('(')
          FunctionParameterListImpl
            FunctionParameterImpl
              LiteralIdentifierImpl
                PsiElement(IDENTIFIER)('b')
              PsiWhiteSpace(' ')
              TypeNameImpl
                LiteralIdentifierImpl
                  PsiElement(IDENTIFIER)('Buffer')
          PsiElement())(')')
          PsiWhiteSpace(' ')
          FunctionResult
            FunctionParameterListImpl
              FunctionParameterImpl
                TypeNameImpl
                  LiteralIdentifierImpl
                    PsiElement(IDENTIFIER)('bool')
        PsiWhiteSpace('\n')
        PsiWhiteSpace('\t')
        FunctionDeclaration(Write)
          LiteralIdentifierImpl
            PsiElement(IDENTIFIER)('Write')
          PsiElement(()('(')
          FunctionParameterListImpl
            FunctionParameterImpl
              LiteralIdentifierImpl
                PsiElement(IDENTIFIER)('b')
              PsiWhiteSpace(' ')
              TypeNameImpl
                LiteralIdentifierImpl
                  PsiElement(IDENTIFIER)('Buffer')
          PsiElement())(')')
          PsiWhiteSpace(' ')
          FunctionResult
            FunctionParameterListImpl
              FunctionParameterImpl
                TypeNameImpl
                  LiteralIdentifierImpl
                    PsiElement(IDENTIFIER)('bool')
        PsiWhiteSpace('\n')
        PsiWhiteSpace('\t')
        FunctionDeclaration(Close)
          LiteralIdentifierImpl
            PsiElement(IDENTIFIER)('Close')
          PsiElement(()('(')
          PsiElement())(')')
        PsiWhiteSpace('\n')
        PsiElement(})('}')

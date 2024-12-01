# Usuário

#### 1. Registro de Usuário:
Como visitante do site, eu quero poder criar uma nova conta informando meus dados básicos, para que eu possa acessar funcionalidades exclusivas para usuários registrados.

##### 1.1 Registro com sucesso
- **Dado** que sou um visitante do site
- **Quando** preencho o formulário com:
  - Nome
  - Email válido
  - Senha
- **Então** devo receber um email de confirmação
  - E ser redirecionado para a página de boas-vindas
  - E ver uma mensagem de "Registro realizado com sucesso"

Cenário 2: Email já cadastrado
Dado que tento me registrar
Quando informo um email já existente no sistema
Então devo ver uma mensagem "Este email já está cadastrado"
E o formulário não deve ser enviado

Cenário 3: Dados inválidos
Dado que tento me registrar
Quando informo dados inválidos (senha fraca/emails diferentes)
Então devo ver mensagens específicas de erro para cada campo
E o formulário não deve ser enviado

- Como usuário, eu quero poder me registrar no sistema, para que eu possa efetuar o login posteriormente
- Como usuário, eu quero poder me autenticar no sistema, para que eu possa acessar a area restrita do sistema
- Como usuário, eu quero poder atualizar os meus dados (nome e senha)
- Como um usuário, eu quero poder resetar minha senha para que eu possa recuperar acesso ao sistema caso esqueça minha senha atual
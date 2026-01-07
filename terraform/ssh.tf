# SSH Key Pair para acesso à instância EC2
# Usa a key pair existente criada na AWS Console

data "aws_key_pair" "existing" {
  key_name = "tc-fiap-customer-key"
}

output "key_pair_name" {
  description = "Nome da key pair utilizada"
  value       = data.aws_key_pair.existing.key_name
}

output "email_searcher_public_dns" {
  value = aws_instance.email-searcher.public_dns
}

output "email_searcher_public_ip" {
  value = aws_instance.email-searcher.public_ip
}

output "application-url" {
  value = "http://${aws_instance.email-searcher.public_dns}:8080"
}

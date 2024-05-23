output "ip_address" {
  value = aws_instance.email-searcher.public_ip
}

output "application-url" {
  value = "http://${aws_instance.email-searcher.public_dns}:8080"
}

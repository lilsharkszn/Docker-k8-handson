output "instance_public_ip" {
  description = "Public IP of Kali EC2 instance"
  value       = module.ec2.public_ip
}

output "instance_id" {
  description = "EC2 instance ID"
  value       = module.ec2.instance_id
}

output "ssh_command" {
  description = "SSH command to connect to the instance"
  value       = "ssh -i ~/container-assessment/Adejare.pem kali@${module.ec2.public_ip}"
}


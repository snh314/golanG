# Domain Health Checker

A simple command-line tool written in Go that checks the email security configuration of a domain by verifying its DNS records (MX, SPF, and DMARC).

## 📖 Overview

Email spoofing is a common attack where malicious actors send emails pretending to be from your domain. To prevent this, domains must be configured with proper DNS records. This tool helps you quickly verify whether a domain has the necessary email security records in place.

The tool checks for:

- **domain** – The domain name that is being checked.
- **MX Record** – Determines if the domain can receive emails.
- **hasSPF** – Specifies which servers are authorized to send emails on behalf of the domain.
- **spfRecord** – The actual SPF record found in the domain's TXT records (if any).
- **hasDMARC** – Indicates whether the domain has a DMARC policy in place.
- **hasdmarcRecord** – The actual DMARC record found in the domain's `_dmarc` TXT records (if any).

## ✨ Features

- 🔍 Look up MX records for a given domain
- 🛡️ Verify SPF (Sender Policy Framework) records
- 🔒 Verify DMARC (Domain-based Message Authentication, Reporting, and Conformance) records
- 📊 Display a clean, readable summary of results
- 💻 Simple CLI interface

## 🚀 Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) (version 1.18 or higher recommended)
- An active internet connection (for DNS lookups)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/snh314/4.email_checker_tool.git
   cd 4.email_checker_tool

   ```

2. Build the project:

   ```bash
   go build -o domain-checker

   ```

3. Run the tool:

   ```bash
   ./domain-checker
   ```

Or run directly with go run:

```bash
go run main.go
```

### 🖥️ Usage

```
Run the program.
When prompted, enter a domain name (e.g., google.com).
The tool will output the results.
```

### Example

```
Input:
Enter a domain: google.com
```

### Output:

```

domain: google.com, hasMX: true, hasSPF: true, spfRecord: v=spf1 include:_spf.google.com ~all, hasDMARC: true, dmarcRecord: v=DMARC1; p=reject; rua=mailto:mailauth-reports@google.com
```

### 🧠 How It Works

The tool uses Go's standard net package to perform DNS lookups:

● MX Lookup – net.LookupMX(domain) fetches the mail exchange records.

● SPF Lookup – net.LookupTXT(domain) fetches TXT records; it scans them for a record starting with v=spf1.

● DMARC Lookup – net.LookupTXT("\_dmarc." + domain) fetches the DMARC record; it scans for a record starting with v=DMARC1.

### ⚠️ Known Bug (Case Sensitivity)

The SPF check currently uses strings.HasPrefix(record, "v=spF1"), which is case-sensitive and incorrect. The correct prefix is v=spf1 (all lowercase). This causes hasSPF to incorrectly report false for domains that actually have a valid SPF record.

### 📚 Why Email Security Records Matter

● Record Purpose

● domain The identity being verified for email security

● MX Tells the internet where to deliver emails for your domain

● hasSPF Indicates whether an SPF policy exists to prevent spoofing

● spfRecord The actual rule listing authorized sending servers

● hasDMARC Shows whether a DMARC policy is enforced

● hasdmarcRecord The actual DMARC rule defining how to handle failed emails

● Without these records, attackers can spoof your domain, damage your reputation, and trick your customers.

### 🛠️ Technologies Used

Language: Go (Golang)

### Standard Library Packages:

● bufio – for reading user input

● fmt – for formatted output

● log – for error logging

● net – for DNS lookups

● os – for standard input

● strings – for string manipulation

#  CRAFTBIT  
[![CI](https://github.com/stackzoo/craftbit/actions/workflows/ci.yaml/badge.svg)](https://github.com/stackzoo/craftbit/actions/workflows/ci.yaml) [![Go Report Card](https://goreportcard.com/badge/github.com/stackzoo/craftbit)](https://goreportcard.com/report/github.com/stackzoo/craftbit)  
[![Latest Release](https://img.shields.io/github/release/stackzoo/craftbit.svg)](https://github.com/stackzoo/craftbit/releases/latest) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)  

<img src="docs/images/logo.png" width="100" height="100" />

<br/>
Bitcoin Swiss Army Knife 🪛 🌕 🔧

## Abstract
This software is a lightweight Command Line Interface (CLI) containing multiple utilities designed for interacting with the Bitcoin ecosystem.  
These utilities are reusable and located within the `pkg` folder.  
Most of these utilities either serve as wrappers for [*btcd*](https://github.com/btcsuite/btcd) libraries or make calls to the [*mempool.space REST APIs*](https://mempool.space/docs/api/rest).  

> [!WARNING] 
> As certain functions leverage the APIs of the public instance of **mempool.space**, this tool may be suboptimal from a privacy perspective and inadvertently expose personally identifiable information (PII) such as transactions or addresses.  


## Example

https://github.com/stackzoo/craftbit/assets/102741679/2ee9dd77-4986-4bcb-ae17-9d94670fc9cc



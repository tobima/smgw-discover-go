# smgw-discover-go
Go libary to discover the IPv6 link local address for the HAN interface of a German Smart Meter Gateway that implements the optional mDNS based auto discovery.

It's known to work with EMH Casa 1.1.

Requires direct connection to the Smart Meter Gateway HAN interface.
**BSI-compliant**: The HAN interface must be operated in a separate network
 isolated from the internet (BSI TR-03109).
**Network setup**: Typically subnet without default route. No IPv4 address configuration required. IPv6 link local addresse are sufficient. 

## Disclaimer

This project is an independent, open-source library and is **not affiliated with, endorsed by, or sponsored by Smart meter vendor**.  

This software is provided **“as is”**, without warranty of any kind, express or implied.  
Use of this library is **at your own risk**.

---

## Data Protection

This library does not collect, store, or transmit data on its own.  
Any processing of metering data, which may be considered personal data under applicable laws, is the responsibility of the integrating application and its operator.

---

## License

This project is licensed under the **MIT License**. See the `LICENSE` file for details.


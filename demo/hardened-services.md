Hardened system services for Windows 10
---

* Services too aggressive to disable (to avoid weird problems with some installers or software)
  - Secondary Logon
  - Windows Image Acquisition
  - Bluetooth Support Service
  - Print Spooler
  - Fax

* Fix Services
  - Credential Manager: Automatic. (to avoid extremely slow logins under certain circumstances)
  - From the Command Prompt (Admin) run: `sc config BITS depend=netprofm` ... This fixes a service dependency that's incorrectly set, and difficult to diagnose

* Not required (unless)
  * using Bluetooth Audio/headphones:
    - AVCTP service (Manual, Trigger Start) For Bluetooth Audio/headphones
    - Bluetooth Audio Gateway Service (Manual, Triggered)
  * using digitally protected video. May affect Netflix or streaming services:
    - Intel(R) Capability Licensing Service TCP IP Interface (Manual)
    - Intel(R) Content Protection HDCP Service (Automatic)
    - Intel(R) Content Protection HECI Service (Manual)
  * using Intel Management Engine for remote management:
    - Intel(R) Dynamic Application Loader Host Interface Service (Automatic, Delayed)
    - Intel(R) Management and Security Application Local Management Service (Automatic, Delayed)
    - Intel(R) TPM Provisioning Service (Automatic)
  * needing Remote Assistance
    - Peer Name Resolution Protocol (Manual)
    - Peer Networking Grouping (Manual)
    - Peer Networking Identity Manager (Manual)
  * using Smart Cards
    - Smart Card (Manual)
    - Smart Card Device Enumeration Service (Manual)
    - Smart Card Removal Policy (Manual)
  * no other anti-malware installed. To disable:
    - download ExecTI.exe from winaero.com/execti/ and run it.
    - start `regedit.exe -m`, change from 3 to 4 at the following key

      ```
      HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Services\SecurityHealthService\Start
      ```
    - on next reboot, Security Center will be disabled.
  * using Xbox
    - Xbox Accessory Management Service (Manual, Trigger Start)
    - Xbox Live Auth Manager (Manual)
    - Xbox Live Game Save (Manual, Trigger Start)
    - Xbox Live Networking Service (Manual)


* Disable Services
  - AllJoyn Router Service: If you don't have IoT clients or want to not use them.
  - Application Layer Gateway Service (Manual): Supports third-party plugins for Internet Connection Sharing
  - Application Management (Manual): Security vulnerability. Only needed if you install/manage software through Group Policy on a domain.
  - BitLocker, Application Management (Manual), Drive Encryption Service (Manual, Trigger Start): If you don't use drive encryption.
  - Certificate Propagation (Manual, Triggered): Only required for Smart Card log on devices.
  - Distributed Link Tracking Client (Automatic): Old service that keeps track of NTFS files as they move around on a server.
  - Downloaded Maps Manager (Automatic, Delayed Start): If you don't use Maps
  - IP Helper (Automatic): If you don't use IPv6.
  - IPsec Policy Agent (Manual, Trigger Start): A really old point to point VPN that nobody should use.
  - Internet Connection Sharing (ICS) (Manual, Trigger Start): Old service, disabling could cause problems with Internet Explorer
  - Offline Files (Manual, Trigger Start): If you connect to a network share which has "Make available offline" enabled; otherwise you can disable it.
  - Microsoft iSCSI Initiator (Manual): If you don't know what an iSCSI target is you don't need it.
  - Netlogon (Manual): If you don't connect to a domain
  - Network Connected Devices Auto-Setup (Manual): If you prefer to manually install network devices, disable this.
  - Parental Controls (Manual): If you don't use parental controls.
  - Program Compatibility Assistant Service (Manual): Enable if you need to run software in Compatibility mode.
  - Remote Access Connection Manager (Automatic): Removes WAN Miniport Adapters.
  - Remote Access Auto Connection Manager (Manual): Not required if you never use RAS.
  - Remote Registry (Disabled): Confirm it's disabled.
  - Routing and Remote Access (Disabled): Confirm it's disabled for Security vulnerability if you don't use RDS.
    - Remote Desktop Services (Manual)
    - Remote Desktop Services UserMode Port Redirection (Manual)
    - Remote Desktop Configuration (Manual)
  - Secure Socket Tunneling Protocol Service (Manual) Disables the WAN Miniport (SSTP), if you don't use this kind of VPN.
  - SNMP Trap (Manual): If you don't know what SNMP traps are you don't need this.
  - SSDP Discovery (Manual): If you don't need to discover UPnP devices on the network—May negatively affect torrent software.
  - UPnP Device Host (Manual) If you don't need to publish UPnP devices from your computer.
  - TCP/IP NetBIOS Helper (Manual, Trigger Start): May cause problems with resolving network hostnames that do not have reverse DNS entries.
  - Touch Keyboard and Handwriting Panel Service (Manual, Trigger Start): If you don't have a touch screen.
  - Retail Demo Service (Manual): Disable.
  - WebClient (Manual): No one really uses WebClient any more. Maybe SharePoint.
  - Windows Error Reporting Service (Manual, Trigger Start): If you will never create or send Windows Error Reports.
  - Windows Insider Service (Manual, Trigger Start): If you will never use this.
  - Windows Mobile Hotspot Service (Manual, Trigger Start): Disable if you won't share internet with others.
  - Windows Remote Management (Manual): Security vulnerability if you don't use it.


* Keep Services
  - Network List Service - Causes problems with network tray icon/discovery
  - Network Location Awareness - Causes problems with network tray icon/discovery
  - Network Location Awareness - Now required by BITS
  - Network List Service - Now required by BITS
  - Server - Now required by Workstation service

* See
  - https://www.deviantart.com/sammilucia/art/Windows-10-20H2-Services-that-are-safe-to-disable-861880643
  - https://www.thewindowsclub.com/which-windows-10-services-safe-to-disable
  - https://www.pcerror-fix.com/windows-10-services-to-disable-for-gaming

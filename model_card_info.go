/*
Decta API

<h2>Introduction</h2><p>This is a documented guide that    describes the process of DECTA card and token operations for both private and business clients. As a new or already    experienced API user you will find in this documentation all necessary and required information to start using DECTA    API and improve your experience with DECTA.</p><h2>Authentication and Authorisation</h2><p>Each request to DECTA API must be signed with a certificate that allows DECTA to identify API clients. This    guarantees the authenticity of the request received from the client.</p><h3>Authentication</h3><p>The authentication process is executed only once for each client.<br>    On successful authentication DECTA assigns a certificate to the client which must be used in every API    request in the future.</p><p>    To access a detailed description of the authentication process, choose <b>Service</b> specification at the top right    corner or click <a        href=\"/?urls.primaryName=SERVICE\" target=\"_blank\">here</a>.</p><h3>Authorisation</h3><p>    Each request must be signed using Jose specifications in particular - <a        href=\"https://tools.ietf.org/html/rfc7515\" target=\"_blank\">RFC7515</a>. It is vital for users to follow the    provided standard and use <b>Base64url    Encoding without Padding</b> for token encoding.</p><h3>Request with filters</h3><p>    Several requests support data filters. Filter syntax can be found below -<p>    <code>        &lt;requestFilter&gt; ::= &lt;filterName&gt;=[eq|ne|gt|ge|lt|le]:&lt;filterValue&gt;    </code></p><ul>Where:    <li>eq - equals</li>    <li>ne - not equals</li>    <li>gt - greater than</li>    <li>ge - greater or equals than</li>    <li>lt - less than</li>    <li>le - less or equals than</li></ul><ul>Example:    <li>cardState: \"eq:BLOCKED_BY_HOLDER\"</li></ul><h2>Changelog</h2>Information about the latest DECTA API changes and improvements can be found <a href=\"/changelog.html\" target=\"_blank\">here.</a>.<h4>Request examples</h4><p>    We recommend users follow the instructions in the next sections below and create the application step by step.</p><p>    To understand the process of DECTA API requests better, we have prepared two basic API calls as examples.</p><p>    <b>POST /v1/api/cards/order</b> request</p><p>For example:</p><ul>    <li>        <b>uri</b> - <br>/v1/api/cards/order    </li>    <li>        <b>body</b> - <br>        {\"card\":{\"accountOwnerRelation\":\"OWNER\",\"currencies\":[\"EUR\"],\"deliveryAddress\":{\"city\":\"Madona\",\"country\":\"LVA\",\"name\":\"Test\",\"phone\":9113,\"shipment\":\"STANDARD\",\"street\":\"62        Madona        street\",\"surname\":\"TestS\",\"zipCode\":\"LV-1000\"},\"holder\":{\"currentAddress\":{\"city\":\"Riga\",\"country\":\"LVA\",\"street\":\"7        Street\",\"zipCode\":\"LV0000\"},\"document\":{\"birthDate\":\"1950-12-31\",\"countryCode\":\"EUR\",\"expiryDate\":\"2056-12-31\",\"issuingDate\":\"2016-12-31\",\"number\":\"TTHolder\",\"subtype\":\"RESIDENT_PERMIT_ID\",\"type\":\"DRIVING_LICENSE\"},\"email\":\"oholder@decta.com\",\"language\":\"EN\",\"maidenName\":\"TestM\",\"mobilePhone\":\"12345678\",\"name\":\"TestS\",\"role\":\"OWNER\",\"surname\":\"TestS\",\"userDefinedField1\":\"2\"},\"passphrase\":\"test2pass\",\"priority\":\"URGENT\",\"productCode\":\"701\",\"supplementary\":\"false\"},\"externalId\":\"Private_New\",\"privateClient\":{\"comment\":\"Agreement        num. 123143\",\"currentAddress\":{\"city\":\"RigaCurrent\",\"country\":\"LVA\",\"months\":4,\"street\":\"6 street        street\",\"zipCode\":\"WScur1009\"},\"document\":{\"birthDate\":\"1950-12-31\",\"countryCode\":\"LVA\",\"expiryDate\":\"2099-12-31\",\"issuingDate\":\"1971-12-31\",\"number\":\"1234\",\"subtype\":\"NATIONAL_ID\",\"type\":\"ID_CARD\"},\"email\":\"email1234\",\"language\":\"en\",\"maidenName\":\"Test\",\"mobilePhone\":\"37128866694\",\"name\":\"Test\",\"pep\":\"NO\",\"surname\":\"PrivateBaseNew\",\"userDefinedField1\":\"2000666\"}}    </li></ul><p><h5>Suppose we have:</h5><ul>    <li><b>Private key</b> -<br>        308204BE020100300D06092A864886F70D0101010500048204A8308204A40201000282010100B73748751B973456931B380163CE168BF884B26B1D27FA6978ACCEF829DF2DBA5D2845ABE997E82531FFD1C22D5E9F1ED6672F6D75D6572B688AE24AB9ED5138D5B5AA484FA04CEC393DB84C833BDDF7507BA1F5DC59030BE6E0FAC16613FADF3D06123621C11FB8F64FE777B9D0308DCEDCBAF78EB76FDAA4151C96E48B68507FA254AEFE44F0C0D103CAFE2A4ADB676E2B4FE26DB4AD671B8F02E74FE4B7E36156F44AA7CCE39DB0879926CEE2EF43502C7F99A1A3DDE72ECAC39E83F7B6819F2C1B8606537D6A04A1D6CC1FF0044F00360EA2F899BD9D9C6D748B20E2EB17CF908D6117B6B5B97490F5D68DC2D039B11EB2D920B8CEE9EFF10D0408A69971020301000102820100584078D67003D5621E59EE103D52F7E9543C16F521863048BDA7FFC9E9E166D56E2A07E5570FA4F2C5B2C40714738F91FB1498F3D0DEFDFE1ACD4B53535BCCD3E39D2DF8C6E2202B692721AF39478D13A3E0E992D420CC26DF7F5F49E9319696117EEB26C7FB8E9C398923A5B80B6057EE5CC4729A7C2194DF948BC0E3358378F1E3381BB62FE98C9372BB6EBBC24E8670E80EF991173D6D6347ADF1F8C9B67920D88072978D025F16F50273EDA0F0AA6FABC6F9687259A21142FD1E5C3DE99EE949DF1E81C7E7CDA4DD806A5CE2105DDC0CF7A2274913971C234D04D7C577C76F2FAE75379A0582FFA0333A01A7BE29ECA203AD7E3E71271EDEFDC0069F684102818100E7C26736089B4FB37BE4873EB2A726BBCD9CA3A3F3EB6BAD1433127614ABC7E6763D118A4A8DB45BBE31A67A853A44BA21EDF50533ACD348A0A9C1F648EE510426ECC845C01185C6793E9D4866281B9C01AE8BA96D0897BCBB0F638696CD6935794506AA3EC14BAB1A7F68DDF70DA21C70249FAD3FE88AB4C6708691327E4BC302818100CA611446E3A9A1997D175BDEADDE02B5C09E3E76024741F2DAAEF8689CF5DCA1C5FA9AE8A6F2F09E837E9D181C5607374DFBAA53B02C3908DCFA1817DEF1797F64743218B1616B55125E618B45F11003BA94D89E357358BD54A1DA324D893FDD9C0B30428A770F410B061C011CDB541E344B4C01B17CB615A1DCAD1F5EEB96BB02818100E40750C9C75A18E73E053254AC2EEE5B6608B2B18433A4341D65CACA47B864ED0A7537A6DB87E56747114EFDC9CBF507368F0CBF5B82B638056C419D73509881FF528612AAD212CF9F47CE3507DE7A9BDAC3C442A5370924F6E0434A8F61F81C56FF65796859837C0C8C43BFF16E868C78827061643A070FAB17D82F508117450281804FE4A32C99138E4819A9EF0AA978CB7914E163A7129F2ED9C09AF255DA20F548A7EF96D7E190668D2D3BFEA856076031E50744E664D6106DFF4E7BD4709EC368173007D6D7AFADBF97D0CA9140BB39A73F312392D16707D13667EECB8CF071D5FA94302914A08BD5119507D9289B2D49FF3AFA7670AADAF70F3F1ED9138FCDA102818100E69108E3E83702F3258503C2FD0E9A34E76750CCE2A20AFDD58F3939351979ECC66373A2151AD313710B2EAFA192DF9F364791FA8473CE1CAE434A044AA8F557F99D756DF4373B5A3F2202ED299979F3C5B4AC1B93ABDB4157ED3901B723B671749DE9279D95759F122400DD5BF3B28F6A8A37383ADFAD9E5C5D49192B20EB43    </li>    <li>Signed by Decta certificate <b>SHA256 Thumbprint</b> base64Url encoded -<br>        WOQEEyrcHaVobHMPORmuMe9iAC5gSI0TpXyIfA1lKI0    </li></ul><h5>JWS preparing: </h5><ul>    <li>        <b>JWS Header</b> -<br> {\"x5t#S256\":\"WOQEEyrcHaVobHMPORmuMe9iAC5gSI0TpXyIfA1lKI0\",\"alg\":\"RS256\"}    </li>    <li><b>JWS Payload</b>-<br>        /v1/api/cards/order{\"card\":{\"accountOwnerRelation\":\"OWNER\",\"currencies\":[\"EUR\"],\"deliveryAddress\":{\"city\":\"Madona\",\"country\":\"LVA\",\"name\":\"Test\",\"phone\":9113,\"shipment\":\"STANDARD\",\"street\":\"62        Madona        street\",\"surname\":\"TestS\",\"zipCode\":\"LV-1000\"},\"holder\":{\"currentAddress\":{\"city\":\"Riga\",\"country\":\"LVA\",\"street\":\"7        Street\",\"zipCode\":\"LV0000\"},\"document\":{\"birthDate\":\"1950-12-31\",\"countryCode\":\"EUR\",\"expiryDate\":\"2056-12-31\",\"issuingDate\":\"2016-12-31\",\"number\":\"TTHolder\",\"subtype\":\"RESIDENT_PERMIT_ID\",\"type\":\"DRIVING_LICENSE\"},\"email\":\"oholder@decta.com\",\"language\":\"EN\",\"maidenName\":\"TestM\",\"mobilePhone\":\"12345678\",\"name\":\"TestS\",\"role\":\"OWNER\",\"surname\":\"TestS\",\"userDefinedField1\":\"2\"},\"passphrase\":\"test2pass\",\"priority\":\"URGENT\",\"productCode\":\"701\",\"supplementary\":\"false\"},\"externalId\":\"Private_New\",\"privateClient\":{\"comment\":\"Agreement        num. 123143\",\"currentAddress\":{\"city\":\"RigaCurrent\",\"country\":\"LVA\",\"months\":4,\"street\":\"6 street        street\",\"zipCode\":\"WScur1009\"},\"document\":{\"birthDate\":\"1950-12-31\",\"countryCode\":\"LVA\",\"expiryDate\":\"2099-12-31\",\"issuingDate\":\"1971-12-31\",\"number\":\"1234\",\"subtype\":\"NATIONAL_ID\",\"type\":\"ID_CARD\"},\"email\":\"email1234\",\"language\":\"en\",\"maidenName\":\"Test\",\"mobilePhone\":\"37128866694\",\"name\":\"Test\",\"pep\":\"NO\",\"surname\":\"PrivateBaseNew\",\"userDefinedField1\":\"2000666\"}}    </li></ul><h5>JWS signing result:</h5><ul>    <li><b>Token-Header</b> (thumbprint encoded using Base64url Encoding without Padding) -<br>        eyJhbGciOiJSUzI1NiIsIng1dCNTMjU2IjoiV09RRUV5cmNIYVZvYkhNUE9SbXVNZTlpQUM1Z1NJMFRwWHlJZkExbEtJMCJ9    </li>    <li style=\"word-wrap: normal\"><b>Token-body</b> (payload encoded using Base64url Encoding without Padding) -<br>        eyJhbGciOiJSUzI1NiIsIng1dCNTMjU2IjoiV09RRUV5cmNIYVZvYkhNUE9SbXVNZTlpQUM1Z1NJMFRwWHlJZkExbEtJMCJ9.L3YxL2FwaS9jYXJkcy9vcmRlcnsiY2FyZCI6eyJhY2NvdW50T3duZXJSZWxhdGlvbiI6Ik9XTkVSIiwiY3VycmVuY2llcyI6WyJFVVIiXSwiZGVsaXZlcnlBZGRyZXNzIjp7ImNpdHkiOiJNYWRvbmEiLCJjb3VudHJ5IjoiTFZBIiwibmFtZSI6IlRlc3QiLCJwaG9uZSI6OTExMywic2hpcG1lbnQiOiJTVEFOREFSRCIsInN0cmVldCI6IjYyIE1hZG9uYSBzdHJlZXQiLCJzdXJuYW1lIjoiVGVzdFMiLCJ6aXBDb2RlIjoiTFYtMTAwMCJ9LCJob2xkZXIiOnsiY3VycmVudEFkZHJlc3MiOnsiY2l0eSI6IlJpZ2EiLCJjb3VudHJ5IjoiTFZBIiwic3RyZWV0IjoiNyBTdHJlZXQiLCJ6aXBDb2RlIjoiTFYwMDAwIn0sImRvY3VtZW50Ijp7ImJpcnRoRGF0ZSI6IjE5NTAtMTItMzEiLCJjb3VudHJ5Q29kZSI6IkVVUiIsImV4cGlyeURhdGUiOiIyMDU2LTEyLTMxIiwiaXNzdWluZ0RhdGUiOiIyMDE2LTEyLTMxIiwibnVtYmVyIjoiVFRIb2xkZXIiLCJzdWJ0eXBlIjoiUkVTSURFTlRfUEVSTUlUX0lEIiwidHlwZSI6IkRSSVZJTkdfTElDRU5TRSJ9LCJlbWFpbCI6Im9ob2xkZXJAZGVjdGEuY29tIiwibGFuZ3VhZ2UiOiJFTiIsIm1haWRlbk5hbWUiOiJUZXN0TSIsIm1vYmlsZVBob25lIjoiMTIzNDU2NzgiLCJuYW1lIjoiVGVzdFMiLCJyb2xlIjoiT1dORVIiLCJzdXJuYW1lIjoiVGVzdFMiLCJ1c2VyRGVmaW5lZEZpZWxkMSI6IjIifSwicGFzc3BocmFzZSI6InRlc3QycGFzcyIsInByaW9yaXR5IjoiVVJHRU5UIiwicHJvZHVjdENvZGUiOiI3MDEiLCJzdXBwbGVtZW50YXJ5IjoiZmFsc2UifSwiZXh0ZXJuYWxJZCI6IlByaXZhdGVfTmV3IiwicHJpdmF0ZUNsaWVudCI6eyJjb21tZW50IjoiQWdyZWVtZW50IG51bS4gMTIzMTQzIiwiY3VycmVudEFkZHJlc3MiOnsiY2l0eSI6IlJpZ2FDdXJyZW50IiwiY291bnRyeSI6IkxWQSIsIm1vbnRocyI6NCwic3RyZWV0IjoiNiBzdHJlZXQgc3RyZWV0IiwiemlwQ29kZSI6IldTY3VyMTAwOSJ9LCJkb2N1bWVudCI6eyJiaXJ0aERhdGUiOiIxOTUwLTEyLTMxIiwiY291bnRyeUNvZGUiOiJMVkEiLCJleHBpcnlEYXRlIjoiMjA5OS0xMi0zMSIsImlzc3VpbmdEYXRlIjoiMTk3MS0xMi0zMSIsIm51bWJlciI6IjEyMzQiLCJzdWJ0eXBlIjoiTkFUSU9OQUxfSUQiLCJ0eXBlIjoiSURfQ0FSRCJ9LCJlbWFpbCI6ImVtYWlsMTMwIiwibGFuZ3VhZ2UiOiJlbiIsIm1haWRlbk5hbWUiOiJUZXN0IiwibW9iaWxlUGhvbmUiOiIzNzEyODg2NjY5NCIsIm5hbWUiOiJUZXN0IiwicGVwIjoiTk8iLCJzdXJuYW1lIjoiUHJpdmF0ZUJhc2VOZXciLCJ1c2VyRGVmaW5lZEZpZWxkMSI6IjIwMDA2NjYifX0.Ks767mp1eA-e1y27GpSmLBhKLud8e96bgYPUpIIlYq5ybeKZ2-WwJsExIIcrnIixaev0ibZ3RrHueKl_uhdF4bVNLiCo8EU2sDnK16e5Z9NL1lS4KZI9WUZHhCIzYJnljdBbHiISSknjdsw3WcvGk6bAYj93U0WplvFxa5lk6Gn2T0lkzujUhJXSHkbuZ541rhwwyztndFbUW07vuTCQ-z_gZuRz8dRbvSnAl0aZdeNWBj0PH3r1GOMG_SR_hpyrf0oEItFGj1ASdOfQjNH78dcMIzle9qh4NzLX8voFTs3u58ktfdK1HhcScumtqWDfVciE7KCBCEpVbpJB9fbiTQ    </li>    <li><b>token-signature</b> (using encoded token header + encoded JWS Payload separated with a \".\" and signed with        the private key) -<br>        emSdpR8owrjpCt0byLf6OSunFEI8wDSC2jgAaSTLVq8ojNZAJskNCdk3bk-XjZhFo0xxgODGrbSjo3CjvJN3jmnSqt1Leq6r4SM-eXxLwHAv9Nmkj-E45JA2sd6MGovmmaQ1if6K5CIIqEXi6pZoohLLslyDTsdTBOb1AQ8eKArMyOhdz6dxBX8RKSATw5FhnVyeshJAhpB5sFopPk97aPkJffl2PXC7CRo14E85icaFUvWQT1MrRNs0XMRBlwatthSYNQvIHq8ahzh3j0XnomgPOdXBZe3soYUaHJs_ZT0JVq8dADNVYIOHux5weB_1nNIC9iRDTbjXLFZdyhZF6Q    </li></ul>Make a request.<p><b>GET /v1/api/cards</b> request<br>    Important - when including query parameters into any of GET requests, they should be preceded by a \"?\" after the URI    in the JWS Payload and combined with \"&\", if multiple are present.<p>For example:</p><ul>    <li>        <b>uri</b> -<br> /v1/api/cards    </li>    <li>        <b>query parameters</b>:<br>        cardAccount: \"eq:00000000\",<br>        cardName: \"eq:Test Test\",<br>        count: \"1000\",<br>        product: \"eq:999\"    </li>    <li>        <b>Resulting URI and query parameters</b> (token-body):<br>        /v1/api/cards?cardAccount=eq:00000000&cardName=eq:Test Test&count=1000&product=eq:999    </li></ul><p><h5>Suppose we have:</h5><ul>    <li><b>Private key</b> -<br>        308204BE020100300D06092A864886F70D0101010500048204A8308204A40201000282010100B73748751B973456931B380163CE168BF884B26B1D27FA6978ACCEF829DF2DBA5D2845ABE997E82531FFD1C22D5E9F1ED6672F6D75D6572B688AE24AB9ED5138D5B5AA484FA04CEC393DB84C833BDDF7507BA1F5DC59030BE6E0FAC16613FADF3D06123621C11FB8F64FE777B9D0308DCEDCBAF78EB76FDAA4151C96E48B68507FA254AEFE44F0C0D103CAFE2A4ADB676E2B4FE26DB4AD671B8F02E74FE4B7E36156F44AA7CCE39DB0879926CEE2EF43502C7F99A1A3DDE72ECAC39E83F7B6819F2C1B8606537D6A04A1D6CC1FF0044F00360EA2F899BD9D9C6D748B20E2EB17CF908D6117B6B5B97490F5D68DC2D039B11EB2D920B8CEE9EFF10D0408A69971020301000102820100584078D67003D5621E59EE103D52F7E9543C16F521863048BDA7FFC9E9E166D56E2A07E5570FA4F2C5B2C40714738F91FB1498F3D0DEFDFE1ACD4B53535BCCD3E39D2DF8C6E2202B692721AF39478D13A3E0E992D420CC26DF7F5F49E9319696117EEB26C7FB8E9C398923A5B80B6057EE5CC4729A7C2194DF948BC0E3358378F1E3381BB62FE98C9372BB6EBBC24E8670E80EF991173D6D6347ADF1F8C9B67920D88072978D025F16F50273EDA0F0AA6FABC6F9687259A21142FD1E5C3DE99EE949DF1E81C7E7CDA4DD806A5CE2105DDC0CF7A2274913971C234D04D7C577C76F2FAE75379A0582FFA0333A01A7BE29ECA203AD7E3E71271EDEFDC0069F684102818100E7C26736089B4FB37BE4873EB2A726BBCD9CA3A3F3EB6BAD1433127614ABC7E6763D118A4A8DB45BBE31A67A853A44BA21EDF50533ACD348A0A9C1F648EE510426ECC845C01185C6793E9D4866281B9C01AE8BA96D0897BCBB0F638696CD6935794506AA3EC14BAB1A7F68DDF70DA21C70249FAD3FE88AB4C6708691327E4BC302818100CA611446E3A9A1997D175BDEADDE02B5C09E3E76024741F2DAAEF8689CF5DCA1C5FA9AE8A6F2F09E837E9D181C5607374DFBAA53B02C3908DCFA1817DEF1797F64743218B1616B55125E618B45F11003BA94D89E357358BD54A1DA324D893FDD9C0B30428A770F410B061C011CDB541E344B4C01B17CB615A1DCAD1F5EEB96BB02818100E40750C9C75A18E73E053254AC2EEE5B6608B2B18433A4341D65CACA47B864ED0A7537A6DB87E56747114EFDC9CBF507368F0CBF5B82B638056C419D73509881FF528612AAD212CF9F47CE3507DE7A9BDAC3C442A5370924F6E0434A8F61F81C56FF65796859837C0C8C43BFF16E868C78827061643A070FAB17D82F508117450281804FE4A32C99138E4819A9EF0AA978CB7914E163A7129F2ED9C09AF255DA20F548A    </li>    <li>Signed by Decta certificate <b>SHA256 Thumbprint</b> base64Url encoded -<br>        WOQEEyrcHaVobHMPORmuMe9iAC5gSI0TpXyIfA1lKI0    </li></ul><h5>JWS preparing: </h5><ul>    <li>        <b>JWS Header</b> -<br> {\"x5t#S256\":\"WOQEEyrcHaVobHMPORmuMe9iAC5gSI0TpXyIfA1lKI0\",\"alg\":\"RS256\"}    </li>    <li><b>JWS Payload</b> -<br>        /v1/api/cardscardAccount=eq:00000000&cardName=eq:Test Test&count=1000&product=eq:999    </li></ul><h5>JWS signing result:</h5><ul>    <li><b>Token-Header</b> (thumbprint encoded using Base64url Encoding without Padding) -<br>        eyJhbGciOiJSUzI1NiIsIng1dCNTMjU2IjoiV09RRUV5cmNIYVZvYkhNUE9SbXVNZTlpQUM1Z1NJMFRwWHlJZkExbEtJMCJ9    </li>    <li style=\"word-wrap: normal\"><b>Token-body</b> (payload encoded using Base64url Encoding without Padding) -<br>        L3YxL2FwaS9jYXJkc2NhcmRBY2NvdW50PWVxOjAwMDAwMDAwJmNhcmROYW1lPWVxOlRlc3QgVGVzdCZjb3VudD0xMDAwJnByb2R1Y3Q9ZXE6OTk5.A2wnCqLNRQi81pzpVvTKDXI4GTMWZ-fjxEt7usuuVVh7BGcEEvRZ4Giy-mLU1e4jBsH5QR6fGT4jLOuTQ_4ItDChH5IBJNHDGTDBoXNNWYFrZS5bgFdtt7--U9i-jPsGySue_iT87Vy8ReJ147mberWN8S18u630gCamhdMdAwPmxIaD_4l1YSKHt29Z_5bK_-uvMIiLrV2ncqRZ_SCnna-GwCceTSa8IKn4uHqb6qngIJYxucWHhye-h-f26KGdhzoTbwz-2u90zNZ0pB3NESclQnmfdKbj8CzY-8Chr9GrAx4p4JLXh12rVGWUTojfsW2rzVgV_H3aw-bCDKVt1g    </li>    <li><b>token-signature</b> (using encoded token header + encoded JWS Payload separated with a \".\" and signed with        the private key) -<br>        A2wnCqLNRQi81pzpVvTKDXI4GTMWZ-fjxEt7usuuVVh7BGcEEvRZ4Giy-mLU1e4jBsH5QR6fGT4jLOuTQ_4ItDChH5IBJNHDGTDBoXNNWYFrZS5bgFdtt7--U9i-jPsGySue_iT87Vy8ReJ147mberWN8S18u630gCamhdMdAwPmxIaD_4l1YSKHt29Z_5bK_-uvMIiLrV2ncqRZ_SCnna-GwCceTSa8IKn4uHqb6qngIJYxucWHhye-h-f26KGdhzoTbwz-2u90zNZ0pB3NESclQnmfdKbj8CzY-8Chr9GrAx4p4JLXh12rVGWUTojfsW2rzVgV_H3aw-bCDKVt1g    </li></ul>Make a request.<p>    <b>If you have further questions regarding DECTA API, please contact our DECTA Support Team -</b><br></p>

API version: 2.11
Contact: support@decta.com
*/

package decta_api

import (
	"encoding/json"
)

// checks if the CardInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardInfo{}

// CardInfo Card information
type CardInfo struct {
	// id
	Id          *string          `json:"id,omitempty"`
	Client      *ClientInfo      `json:"client,omitempty"`
	CardAccount *CardAccountInfo `json:"cardAccount,omitempty"`
	// Card pseudo-pan
	Ppan *string `json:"ppan,omitempty"`
	// Name on card
	CardName *string `json:"cardName,omitempty"`
	// Product code
	Product *string `json:"product,omitempty"`
	// Card state
	State *string `json:"state,omitempty"`
	// Card expiry date
	Expiry *string `json:"expiry,omitempty"`
	// Design ID for plastic card product provided by Decta
	Design *string `json:"design,omitempty"`
	// Priority flag for plastic card embossing
	Priority *string `json:"priority,omitempty"`
}

// NewCardInfo instantiates a new CardInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardInfo() *CardInfo {
	this := CardInfo{}
	return &this
}

// NewCardInfoWithDefaults instantiates a new CardInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardInfoWithDefaults() *CardInfo {
	this := CardInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CardInfo) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInfo) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CardInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CardInfo) SetId(v string) {
	o.Id = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *CardInfo) GetClient() ClientInfo {
	if o == nil || IsNil(o.Client) {
		var ret ClientInfo
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInfo) GetClientOk() (*ClientInfo, bool) {
	if o == nil || IsNil(o.Client) {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *CardInfo) HasClient() bool {
	if o != nil && !IsNil(o.Client) {
		return true
	}

	return false
}

// SetClient gets a reference to the given ClientInfo and assigns it to the Client field.
func (o *CardInfo) SetClient(v ClientInfo) {
	o.Client = &v
}

// GetCardAccount returns the CardAccount field value if set, zero value otherwise.
func (o *CardInfo) GetCardAccount() CardAccountInfo {
	if o == nil || IsNil(o.CardAccount) {
		var ret CardAccountInfo
		return ret
	}
	return *o.CardAccount
}

// GetCardAccountOk returns a tuple with the CardAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInfo) GetCardAccountOk() (*CardAccountInfo, bool) {
	if o == nil || IsNil(o.CardAccount) {
		return nil, false
	}
	return o.CardAccount, true
}

// HasCardAccount returns a boolean if a field has been set.
func (o *CardInfo) HasCardAccount() bool {
	if o != nil && !IsNil(o.CardAccount) {
		return true
	}

	return false
}

// SetCardAccount gets a reference to the given CardAccountInfo and assigns it to the CardAccount field.
func (o *CardInfo) SetCardAccount(v CardAccountInfo) {
	o.CardAccount = &v
}

// GetPpan returns the Ppan field value if set, zero value otherwise.
func (o *CardInfo) GetPpan() string {
	if o == nil || IsNil(o.Ppan) {
		var ret string
		return ret
	}
	return *o.Ppan
}

// GetPpanOk returns a tuple with the Ppan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInfo) GetPpanOk() (*string, bool) {
	if o == nil || IsNil(o.Ppan) {
		return nil, false
	}
	return o.Ppan, true
}

// HasPpan returns a boolean if a field has been set.
func (o *CardInfo) HasPpan() bool {
	if o != nil && !IsNil(o.Ppan) {
		return true
	}

	return false
}

// SetPpan gets a reference to the given string and assigns it to the Ppan field.
func (o *CardInfo) SetPpan(v string) {
	o.Ppan = &v
}

// GetCardName returns the CardName field value if set, zero value otherwise.
func (o *CardInfo) GetCardName() string {
	if o == nil || IsNil(o.CardName) {
		var ret string
		return ret
	}
	return *o.CardName
}

// GetCardNameOk returns a tuple with the CardName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInfo) GetCardNameOk() (*string, bool) {
	if o == nil || IsNil(o.CardName) {
		return nil, false
	}
	return o.CardName, true
}

// HasCardName returns a boolean if a field has been set.
func (o *CardInfo) HasCardName() bool {
	if o != nil && !IsNil(o.CardName) {
		return true
	}

	return false
}

// SetCardName gets a reference to the given string and assigns it to the CardName field.
func (o *CardInfo) SetCardName(v string) {
	o.CardName = &v
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *CardInfo) GetProduct() string {
	if o == nil || IsNil(o.Product) {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInfo) GetProductOk() (*string, bool) {
	if o == nil || IsNil(o.Product) {
		return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *CardInfo) HasProduct() bool {
	if o != nil && !IsNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *CardInfo) SetProduct(v string) {
	o.Product = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CardInfo) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInfo) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CardInfo) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *CardInfo) SetState(v string) {
	o.State = &v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *CardInfo) GetExpiry() string {
	if o == nil || IsNil(o.Expiry) {
		var ret string
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInfo) GetExpiryOk() (*string, bool) {
	if o == nil || IsNil(o.Expiry) {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *CardInfo) HasExpiry() bool {
	if o != nil && !IsNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given string and assigns it to the Expiry field.
func (o *CardInfo) SetExpiry(v string) {
	o.Expiry = &v
}

// GetDesign returns the Design field value if set, zero value otherwise.
func (o *CardInfo) GetDesign() string {
	if o == nil || IsNil(o.Design) {
		var ret string
		return ret
	}
	return *o.Design
}

// GetDesignOk returns a tuple with the Design field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInfo) GetDesignOk() (*string, bool) {
	if o == nil || IsNil(o.Design) {
		return nil, false
	}
	return o.Design, true
}

// HasDesign returns a boolean if a field has been set.
func (o *CardInfo) HasDesign() bool {
	if o != nil && !IsNil(o.Design) {
		return true
	}

	return false
}

// SetDesign gets a reference to the given string and assigns it to the Design field.
func (o *CardInfo) SetDesign(v string) {
	o.Design = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *CardInfo) GetPriority() string {
	if o == nil || IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardInfo) GetPriorityOk() (*string, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *CardInfo) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *CardInfo) SetPriority(v string) {
	o.Priority = &v
}

func (o CardInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Client) {
		toSerialize["client"] = o.Client
	}
	if !IsNil(o.CardAccount) {
		toSerialize["cardAccount"] = o.CardAccount
	}
	if !IsNil(o.Ppan) {
		toSerialize["ppan"] = o.Ppan
	}
	if !IsNil(o.CardName) {
		toSerialize["cardName"] = o.CardName
	}
	if !IsNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	if !IsNil(o.Design) {
		toSerialize["design"] = o.Design
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	return toSerialize, nil
}

type NullableCardInfo struct {
	value *CardInfo
	isSet bool
}

func (v NullableCardInfo) Get() *CardInfo {
	return v.value
}

func (v *NullableCardInfo) Set(val *CardInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCardInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCardInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardInfo(val *CardInfo) *NullableCardInfo {
	return &NullableCardInfo{value: val, isSet: true}
}

func (v NullableCardInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

